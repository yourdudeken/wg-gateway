// API Base URL
const API_BASE = '';

// Initialize dashboard
document.addEventListener('DOMContentLoaded', () => {
    loadStatus();
    loadConfig();
    loadPeers();
    loadServices();

    // Refresh data every 10 seconds
    setInterval(() => {
        loadStatus();
        loadPeers();
        loadServices();
    }, 10000);
});

// Load status
async function loadStatus() {
    try {
        const response = await fetch(`${API_BASE}/api/status`);
        const data = await response.json();

        document.getElementById('peerCount').textContent = data.peer_count;
        document.getElementById('serviceCount').textContent = data.service_count;
        document.getElementById('vpsIP').textContent = data.vps_ip || '-';
        document.getElementById('readyStatus').textContent = data.ready ? 'Ready' : 'Setup Required';

        const statusBadge = document.getElementById('statusBadge');
        const statusDot = statusBadge.querySelector('.status-dot');
        const statusText = statusBadge.querySelector('.status-text');

        if (data.ready) {
            statusDot.style.background = 'var(--success)';
            statusText.textContent = 'Production Ready';
        } else {
            statusDot.style.background = 'var(--warning)';
            statusText.textContent = 'Configuration Required';
        }
    } catch (error) {
        console.error('Error loading status:', error);
        showNotification('Failed to load status', 'error');
    }
}

// Load configuration
async function loadConfig() {
    try {
        const response = await fetch(`${API_BASE}/api/config`);
        const data = await response.json();

        document.getElementById('vpsIpInput').value = data.vps.ip || '';
        document.getElementById('vpsUserInput').value = data.vps.ssh_user || '';
        document.getElementById('proxyEmailInput').value = data.proxy.email || '';
        document.getElementById('projectInput').value = data.project || '';
    } catch (error) {
        console.error('Error loading config:', error);
    }
}

// Update configuration
async function updateConfig(event) {
    event.preventDefault();

    const updates = [
        { key: 'vps.ip', value: document.getElementById('vpsIpInput').value },
        { key: 'vps.user', value: document.getElementById('vpsUserInput').value },
        { key: 'proxy.email', value: document.getElementById('proxyEmailInput').value },
        { key: 'project', value: document.getElementById('projectInput').value }
    ];

    try {
        for (const update of updates) {
            await fetch(`${API_BASE}/api/config/update`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(update)
            });
        }

        showNotification('Configuration updated successfully', 'success');
        loadStatus();
    } catch (error) {
        console.error('Error updating config:', error);
        showNotification('Failed to update configuration', 'error');
    }
}

// Load peers
async function loadPeers() {
    try {
        const response = await fetch(`${API_BASE}/api/peers`);
        const peers = await response.json();

        const container = document.getElementById('peersList');

        if (!peers || peers.length === 0) {
            container.innerHTML = '<p class="empty-state">No peers configured</p>';
            return;
        }

        container.innerHTML = peers.map(peer => `
            <div class="list-item">
                <div class="list-item-content">
                    <h4>${peer.name}</h4>
                    <p>IP: ${peer.wg_ip} | Keepalive: ${peer.keepalive}s</p>
                </div>
                <div class="list-item-actions">
                    <button class="btn btn-secondary btn-sm" onclick="copyPublicKey('${peer.public_key}')">
                        Copy Key
                    </button>
                </div>
            </div>
        `).join('');

        // Update peer select in service modal
        updatePeerSelect(peers);
    } catch (error) {
        console.error('Error loading peers:', error);
    }
}

// Update peer select dropdown
function updatePeerSelect(peers) {
    const select = document.getElementById('servicePeer');
    select.innerHTML = '<option value="">Select a peer...</option>' +
        peers.map(peer => `<option value="${peer.name}">${peer.name} (${peer.wg_ip})</option>`).join('');
}

// Load services
async function loadServices() {
    try {
        const response = await fetch(`${API_BASE}/api/services`);
        const services = await response.json();

        const container = document.getElementById('servicesList');

        if (!services || services.length === 0) {
            container.innerHTML = '<p class="empty-state">No services configured</p>';
            return;
        }

        container.innerHTML = services.map(service => `
            <div class="list-item">
                <div class="list-item-content">
                    <h4>${service.domain}</h4>
                    <p>Port: ${service.port} | Peer: ${service.peer_name}</p>
                </div>
                <div class="list-item-actions">
                    <button class="btn btn-danger btn-sm" onclick="deleteService('${service.domain}')">
                        Delete
                    </button>
                </div>
            </div>
        `).join('');
    } catch (error) {
        console.error('Error loading services:', error);
    }
}

// Show add peer modal
function showAddPeerModal() {
    document.getElementById('addPeerModal').classList.add('active');
}

// Show add service modal
function showAddServiceModal() {
    document.getElementById('addServiceModal').classList.add('active');
}

// Close modal
function closeModal(modalId) {
    document.getElementById(modalId).classList.remove('active');
}

// Add peer
async function addPeer(event) {
    event.preventDefault();

    const name = document.getElementById('peerName').value;
    const ip = document.getElementById('peerIP').value;

    try {
        const response = await fetch(`${API_BASE}/api/peers/add`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, ip })
        });

        if (response.ok) {
            showNotification(`Peer ${name} added successfully`, 'success');
            closeModal('addPeerModal');
            document.getElementById('peerName').value = '';
            document.getElementById('peerIP').value = '';
            loadPeers();
            loadStatus();
        } else {
            const error = await response.text();
            showNotification(`Failed to add peer: ${error}`, 'error');
        }
    } catch (error) {
        console.error('Error adding peer:', error);
        showNotification('Failed to add peer', 'error');
    }
}

// Add service
async function addService(event) {
    event.preventDefault();

    const domain = document.getElementById('serviceDomain').value;
    const port = parseInt(document.getElementById('servicePort').value);
    const peer_name = document.getElementById('servicePeer').value;

    try {
        const response = await fetch(`${API_BASE}/api/services/add`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ domain, port, peer_name })
        });

        if (response.ok) {
            showNotification(`Service ${domain} added successfully`, 'success');
            closeModal('addServiceModal');
            document.getElementById('serviceDomain').value = '';
            document.getElementById('servicePort').value = '';
            document.getElementById('servicePeer').value = '';
            loadServices();
            loadStatus();
        } else {
            const error = await response.text();
            showNotification(`Failed to add service: ${error}`, 'error');
        }
    } catch (error) {
        console.error('Error adding service:', error);
        showNotification('Failed to add service', 'error');
    }
}

// Delete service
async function deleteService(domain) {
    if (!confirm(`Are you sure you want to delete service ${domain}?`)) {
        return;
    }

    try {
        const response = await fetch(`${API_BASE}/api/services/delete?domain=${encodeURIComponent(domain)}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            showNotification(`Service ${domain} deleted successfully`, 'success');
            loadServices();
            loadStatus();
        } else {
            const error = await response.text();
            showNotification(`Failed to delete service: ${error}`, 'error');
        }
    } catch (error) {
        console.error('Error deleting service:', error);
        showNotification('Failed to delete service', 'error');
    }
}

// Copy public key to clipboard
function copyPublicKey(key) {
    navigator.clipboard.writeText(key).then(() => {
        showNotification('Public key copied to clipboard', 'success');
    }).catch(err => {
        console.error('Failed to copy:', err);
        showNotification('Failed to copy key', 'error');
    });
}

// Show notification
function showNotification(message, type = 'info') {
    const notification = document.createElement('div');
    notification.className = `notification notification-${type}`;
    notification.textContent = message;
    notification.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        padding: 1rem 1.5rem;
        background: ${type === 'success' ? 'var(--success)' : type === 'error' ? 'var(--danger)' : 'var(--primary)'};
        color: white;
        border-radius: 8px;
        box-shadow: 0 4px 12px var(--shadow);
        z-index: 2000;
        animation: slideIn 0.3s ease;
    `;

    document.body.appendChild(notification);

    setTimeout(() => {
        notification.style.animation = 'slideOut 0.3s ease';
        setTimeout(() => notification.remove(), 300);
    }, 3000);
}

// Add animation styles
const style = document.createElement('style');
style.textContent = `
    @keyframes slideIn {
        from {
            transform: translateX(400px);
            opacity: 0;
        }
        to {
            transform: translateX(0);
            opacity: 1;
        }
    }
    
    @keyframes slideOut {
        from {
            transform: translateX(0);
            opacity: 1;
        }
        to {
            transform: translateX(400px);
            opacity: 0;
        }
    }
    
    .btn-sm {
        padding: 0.5rem 1rem;
        font-size: 0.85rem;
    }
`;
document.head.appendChild(style);
