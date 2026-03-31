let currentSpec = null;
let selectedElement = null;
let currentEndpointBaseUrl = "";

function getAppConfig() {
    return window.API_CONSOLE_CONFIG || {};
}

function normalizeBase(value) {
    return (value || "").replace(/\/+$/, "");
}

function getApiOrigin() {
    const configuredOrigin = normalizeBase(getAppConfig().apiOrigin);
    if (configuredOrigin) {
        return configuredOrigin;
    }
    if (window.location.origin && window.location.origin !== "null") {
        return window.location.origin;
    }
    return "http://localhost:8000";
}

function getConfiguredApiBasePath() {
    return getAppConfig().apiBasePath || "";
}

function buildApiUrl(path = "") {
    return `${getApiOrigin()}${path}`;
}

async function fetchSwaggerSpec() {
    const configuredPaths = getAppConfig().swaggerSpecPaths;
    const candidates = Array.isArray(configuredPaths) && configuredPaths.length
        ? configuredPaths.map((path) => buildApiUrl(path))
        : [
            buildApiUrl('/api/v1/docs/doc.json'),
            buildApiUrl('/docs/swagger.json'),
            buildApiUrl('/swagger/doc.json')
        ];

    let lastError = null;

    for (const url of candidates) {
        try {
            const res = await fetch(url, {
                headers: { 'Accept': 'application/json' }
            });

            const raw = await res.text();

            if (!res.ok) {
                throw new Error(`HTTP ${res.status}`);
            }

            try {
                return JSON.parse(raw);
            } catch (parseErr) {
                const preview = raw.trim().slice(0, 120);
                throw new Error(`Invalid JSON from ${url}: ${preview || parseErr.message}`);
            }
        } catch (err) {
            lastError = err;
        }
    }

    throw lastError || new Error('Swagger spec not found');
}

function toggleTheme() {
    const b = document.body;
    const isDark = b.getAttribute('data-theme') === 'dark';
    b.setAttribute('data-theme', isDark ? 'light' : 'dark');
    document.getElementById('themeToggle').className = isDark ? 'bx bx-sun' : 'bx bx-moon';
}

async function loadSwagger() {
    try {
        currentSpec = await fetchSwaggerSpec();

        // Fix Title and Version
        document.getElementById('dynamicSidebarTitle').textContent = (currentSpec.info.title || "API").toUpperCase();
        document.getElementById('dynamicApiVersion').textContent = "v" + (currentSpec.info.version || "1.0");

        renderSidebarFolders();
    } catch (err) {
        document.getElementById('sidebar-list').innerHTML = `<div style="padding:20px; font-size:11px; color:var(--error)">Failed to load Swagger spec.<br>${err.message}</div>`;
    }
}

function renderSidebarFolders() {
    const container = document.getElementById('sidebar-list');
    container.innerHTML = '';
    const groups = {};

    Object.keys(currentSpec.paths).forEach(path => {
        const parts = path.split('/').filter(p => p);
        let groupName = 'GENERAL';
        if (parts.length > 0) {
            let idx = parts[0].toLowerCase() === 'api' ? 1 : 0;
            if (parts[idx] && /^v\d+$/i.test(parts[idx])) idx++;
            groupName = parts[idx] ? parts[idx].toUpperCase() : 'GENERAL';
        }
        if (!groups[groupName]) groups[groupName] = [];
        Object.keys(currentSpec.paths[path]).forEach(method => groups[groupName].push({ path, method }));
    });

    Object.keys(groups).sort().forEach(groupName => {
        const groupDiv = document.createElement('div');
        groupDiv.className = 'folder-group';
        groupDiv.innerHTML = `
                    <div class="folder-header" onclick="this.parentElement.classList.toggle('open')">
                        <span><i class='bx bxs-folder' style="margin-right:8px; color:var(--accent)"></i> ${groupName}</span>
                        <i class='bx bx-chevron-right folder-icon'></i>
                    </div>
                    <div class="folder-content"></div>
                `;
        const contentDiv = groupDiv.querySelector('.folder-content');
        groups[groupName].forEach(api => {
            const item = document.createElement('div');
            item.className = 'endpoint-item';
            const color = api.method === 'get' ? 'var(--success)' : api.method === 'post' ? 'var(--warning)' : 'var(--error)';
            const endpoint = currentSpec.paths[api.path][api.method];
            const apiName = endpoint.summary || endpoint.description || api.path;
            item.innerHTML = `
                        <span class="method-tag" style="color:${color}">${api.method.toUpperCase()}</span>
                        <div class="endpoint-meta">
                            <span class="endpoint-title">${apiName}</span>
                            <span class="endpoint-path">${api.path}</span>
                        </div>
                    `;
            item.onclick = (e) => { e.stopPropagation(); selectEndpoint(item, api.path, api.method); };
            contentDiv.appendChild(item);
        });
        container.appendChild(groupDiv);
    });
}

function selectEndpoint(el, path, method) {
    if (selectedElement) selectedElement.classList.remove('selected');
    el.classList.add('selected'); selectedElement = el;

    const basePath = getConfiguredApiBasePath() || currentSpec.basePath || '';
    currentEndpointBaseUrl = buildApiUrl(`${basePath}${path}`);
    document.getElementById('urlInput').value = currentEndpointBaseUrl;
    document.getElementById('methodSelect').value = method.toUpperCase();

    const endpoint = currentSpec.paths[path][method];
    renderEndpointDocs(path, method, endpoint);
    populateQueryParamRows(endpoint);
    populateHeaderRows(endpoint);
    refreshUrlFromParams();

    // Fix Body Logic
    let bodyJSON = {};
    const bodyParam = (endpoint.parameters || []).find(p => p.in === 'body');
    if (bodyParam && bodyParam.schema) {
        bodyJSON = parseSchema(bodyParam.schema);
    }
    document.getElementById('bodyEditor').value = JSON.stringify(bodyJSON, null, 2);
}

function parseSchema(schema) {
    let ref = schema.$ref || (schema.items && schema.items.$ref);
    if (ref) {
        const refKey = ref.replace('#/definitions/', '');
        const definition = currentSpec.definitions[refKey];
        if (!definition) return {};

        let obj = {};
        if (definition.properties) {
            Object.keys(definition.properties).forEach(key => {
                const prop = definition.properties[key];
                if (prop.$ref) obj[key] = parseSchema(prop);
                else if (prop.type === 'string') obj[key] = "";
                else if (prop.type === 'boolean') obj[key] = false;
                else if (prop.type === 'integer' || prop.type === 'number') obj[key] = 0;
                else if (prop.type === 'array') obj[key] = [];
                else obj[key] = null;
            });
        }
        return obj;
    }
    return {};
}

function renderEndpointDocs(path, method, endpoint) {
    const descriptionEl = document.getElementById('endpointDescription');
    const summary = endpoint.summary || 'Untitled API';
    const description = endpoint.description || 'No description available.';
    const tag = endpoint.tags && endpoint.tags.length ? endpoint.tags[0] : 'General';
    const fullPath = `${currentSpec.basePath || ''}${path}`;
    const parameters = endpoint.parameters || [];
    const groupedParams = parameters.reduce((acc, param) => {
        const key = param.in || 'unknown';
        if (!acc[key]) acc[key] = [];
        acc[key].push(param);
        return acc;
    }, {});

    const paramsMarkup = parameters.length
        ? Object.keys(groupedParams).map(groupName => `
                <div class="docs-section-label">${groupName} Params</div>
                <div class="params-list">
                    ${groupedParams[groupName].map(param => `
                        <div class="param-card">
                            <div class="param-name">${param.name}</div>
                            <div class="param-meta">${param.in || 'unknown'} | ${param.type || (param.schema ? 'object' : 'any')} | ${param.required ? 'required' : 'optional'}</div>
                            <div class="param-desc">${param.description || 'No description available.'}</div>
                        </div>
                    `).join('')}
                </div>
            `).join('')
        : `<div class="param-card"><div class="param-desc">No parameters documented for this endpoint.</div></div>`;

    descriptionEl.innerHTML = `
                <h2 class="docs-title">${summary}</h2>
                <div class="docs-meta">
                    <span class="docs-chip">${method.toUpperCase()}</span>
                    <span class="docs-chip">${tag}</span>
                </div>
                <div class="docs-section-label">API Name</div>
                <p class="docs-description">${summary}</p>
                <div class="docs-section-label">Description</div>
                <p class="docs-description">${description}</p>
                <div class="docs-section-label">Path</div>
                <p class="docs-path">${fullPath}</p>
                <div class="docs-section-label">Parameters</div>
                ${paramsMarkup}
            `;
}

function createHeaderRow() {
    const row = document.createElement('tr');
    row.className = 'header-row';
    row.innerHTML = `
                <td class="header-row-toggle"><input class="header-enable" type="checkbox" checked></td>
                <td class="header-row-actions"><button type="button" class="header-remove-btn" onclick="removeHeaderRow(this)" title="Remove row">&times;</button></td>
                <td><input type="text" class="header-key" placeholder="Key"></td>
                <td><input type="text" class="header-value" placeholder="Value"></td>
                <td><input type="text" class="header-desc-input" placeholder="Description"></td>
            `;
    return row;
}

function createQueryParamRow() {
    const row = document.createElement('tr');
    row.className = 'query-param-row';
    row.innerHTML = `
                <td class="param-row-toggle"><input class="param-enable" type="checkbox" checked></td>
                <td class="param-row-actions"><button type="button" class="param-remove-btn" onclick="removeQueryParamRow(this)" title="Remove row">&times;</button></td>
                <td><input type="text" class="param-key" placeholder="Key"></td>
                <td><input type="text" class="param-value" placeholder="Value"></td>
                <td><input type="text" class="param-desc-input" placeholder="Description"></td>
            `;
    return row;
}

function createQueryParamRowWithData(param = {}) {
    const row = createQueryParamRow();
    row.querySelector('.param-enable').checked = !!param.enabled;
    row.querySelector('.param-key').value = param.key || '';
    row.querySelector('.param-value').value = param.value || '';
    row.querySelector('.param-desc-input').value = param.description || '';
    return row;
}

function createHeaderRowWithData(header = {}) {
    const row = createHeaderRow();
    row.querySelector('.header-enable').checked = !!header.enabled;
    row.querySelector('.header-key').value = header.key || '';
    row.querySelector('.header-value').value = header.value || '';
    row.querySelector('.header-desc-input').value = header.description || '';
    return row;
}

function addHeaderRow() {
    const row = createHeaderRow();
    document.getElementById('headersTableBody').appendChild(row);
    row.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
    row.querySelector('.header-key')?.focus();
}

function addQueryParamRow() {
    const row = createQueryParamRow();
    document.getElementById('queryParamsTableBody').appendChild(row);
    row.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
    row.querySelector('.param-key')?.focus();
    refreshUrlFromParams();
}

function removeHeaderRow(button) {
    const tableBody = document.getElementById('headersTableBody');
    const rows = tableBody.querySelectorAll('.header-row');

    if (rows.length === 1) {
        const row = rows[0];
        row.querySelector('.header-enable').checked = false;
        row.querySelector('.header-key').value = '';
        row.querySelector('.header-value').value = '';
        row.querySelector('.header-desc-input').value = '';
        return;
    }

    button.closest('.header-row')?.remove();
}

function removeQueryParamRow(button) {
    const tableBody = document.getElementById('queryParamsTableBody');
    const rows = tableBody.querySelectorAll('.query-param-row');

    if (rows.length === 1) {
        const row = rows[0];
        row.querySelector('.param-enable').checked = false;
        row.querySelector('.param-key').value = '';
        row.querySelector('.param-value').value = '';
        row.querySelector('.param-desc-input').value = '';
        refreshUrlFromParams();
        return;
    }

    button.closest('.query-param-row')?.remove();
    refreshUrlFromParams();
}

function getSwaggerQueryParams(endpoint) {
    const params = endpoint.parameters || [];

    return params
        .filter(param => param.in === 'query')
        .map(param => ({
            key: param.name || '',
            value: '',
            description: param.description || '',
            enabled: !!param.required
        }));
}

function populateQueryParamRows(endpoint) {
    const tableBody = document.getElementById('queryParamsTableBody');
    tableBody.innerHTML = '';

    const swaggerQueryParams = getSwaggerQueryParams(endpoint);

    if (swaggerQueryParams.length) {
        swaggerQueryParams.forEach(param => {
            tableBody.appendChild(createQueryParamRowWithData(param));
        });
    } else {
        tableBody.appendChild(createQueryParamRow());
    }
    refreshUrlFromParams();
}

function getSwaggerHeaderParams(endpoint) {
    const headers = [];
    const params = endpoint.parameters || [];

    params
        .filter(param => param.in === 'header')
        .forEach(param => {
            headers.push({
                key: param.name || '',
                value: '',
                description: param.description || '',
                enabled: !!param.required
            });
        });

    const hasBearerSecurity = Array.isArray(endpoint.security) &&
        endpoint.security.some(item => item.BearerAuth !== undefined);

    const hasAuthorizationHeader = headers.some(header =>
        header.key.toLowerCase() === 'authorization'
    );

    if (hasBearerSecurity && !hasAuthorizationHeader) {
        headers.unshift({
            key: 'Authorization',
            value: '',
            description: 'Bearer token',
            enabled: false
        });
    }

    return headers;
}

function populateHeaderRows(endpoint) {
    const tableBody = document.getElementById('headersTableBody');
    tableBody.innerHTML = '';

    const swaggerHeaders = getSwaggerHeaderParams(endpoint);

    if (swaggerHeaders.length) {
        swaggerHeaders.forEach(header => {
            tableBody.appendChild(createHeaderRowWithData(header));
        });
    } else {
        tableBody.appendChild(createHeaderRow());
    }
}

function getCustomHeaders() {
    const headers = {};
    const rows = document.querySelectorAll('#headersTableBody .header-row');

    rows.forEach(row => {
        const enabled = row.querySelector('.header-enable')?.checked;
        const key = row.querySelector('.header-key')?.value?.trim();
        const value = row.querySelector('.header-value')?.value ?? '';

        if (enabled && key) {
            headers[key] = value;
        }
    });

    return headers;
}

function applyQueryParamsToUrl(url) {
    const finalUrl = new URL(url);
    const rows = document.querySelectorAll('#queryParamsTableBody .query-param-row');

    rows.forEach(row => {
        const enabled = row.querySelector('.param-enable')?.checked;
        const key = row.querySelector('.param-key')?.value?.trim();
        const value = row.querySelector('.param-value')?.value ?? '';

        if (enabled && key) {
            finalUrl.searchParams.set(key, value);
        }
    });

    return finalUrl.toString();
}

function refreshUrlFromParams() {
    const urlInput = document.getElementById('urlInput');
    const baseUrl = currentEndpointBaseUrl || urlInput.value.split('?')[0];
    urlInput.value = applyQueryParamsToUrl(baseUrl);
}

async function sendRequest() {
    const url = applyQueryParamsToUrl(document.getElementById('urlInput').value);
    const method = document.getElementById('methodSelect').value;
    const body = document.getElementById('bodyEditor').value;
    const token = document.getElementById('bearerToken').value;
    const viewer = document.getElementById('responseBody');
    viewer.textContent = "// Loading...";
    const start = performance.now();
    try {
        const headers = { 'Content-Type': 'application/json' };
        if (token) headers['Authorization'] = `Bearer ${token}`;
        Object.assign(headers, getCustomHeaders());
        const options = { method, headers };
        if (method !== 'GET') options.body = body;
        const res = await fetch(url, options);
        const raw = await res.text();
        let data;

        try {
            data = raw ? JSON.parse(raw) : {};
        } catch {
            data = raw;
        }

        document.getElementById('statusCode').textContent = res.status;
        document.getElementById('resTime').textContent = Math.round(performance.now() - start) + "ms";
        viewer.textContent = typeof data === 'string' ? data : JSON.stringify(data, null, 2);
    } catch (err) { viewer.textContent = "Error: " + err.message; }
}

function switchTab(n) {
    [0, 1, 2, 3, 4].forEach(i => {
        document.getElementById(`panel-${i}`).style.display = i === n ? 'block' : 'none';
        document.querySelectorAll('.nav-link')[i].classList.toggle('active', i === n);
    });
}

function copyResponse() { navigator.clipboard.writeText(document.getElementById('responseBody').textContent); }
window.onload = () => {
    loadSwagger();

    document.getElementById('queryParamsTableBody').addEventListener('input', (event) => {
        if (event.target.closest('.query-param-row')) {
            refreshUrlFromParams();
        }
    });

    document.getElementById('queryParamsTableBody').addEventListener('change', (event) => {
        if (event.target.closest('.query-param-row')) {
            refreshUrlFromParams();
        }
    });
};
