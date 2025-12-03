import { get, getCurrentUser, checkAuth, logout } from './api.js';

document.addEventListener('DOMContentLoaded', () => {
  checkAuth();

  const user = getCurrentUser();
  if (user) {
    const nameEl = document.getElementById('userName');
    if (nameEl) nameEl.textContent = user.nama || user.name || user.email || '';
  }

  const logoutBtn = document.getElementById('logoutBtn');
  if (logoutBtn) logoutBtn.addEventListener('click', () => logout());

  loadPending();
});

async function loadPending() {
  try {
    const data = await get('/peminjaman/pending');
    const container = document.getElementById('pendingList');

    if (!container) return;

    if (data && Array.isArray(data) && data.length > 0) {
      container.innerHTML = data
        .map(
          (p) => `
        <div class="border-l-4 border-yellow-500 p-4 bg-gray-50">
          <div class="flex justify-between">
            <div>
              <p class="font-semibold">${p.ruangan ? p.ruangan.nama_ruangan : 'Barang'}</p>
              <p class="text-sm text-gray-600">
                ${new Date(p.tanggal_mulai).toLocaleDateString('id-ID')} - 
                ${new Date(p.tanggal_selesai).toLocaleDateString('id-ID')}
              </p>
              <p class="text-sm text-gray-600">Peminjam: ${p.peminjam ? p.peminjam.nama : 'N/A'}</p>
            </div>
            <a href="verifikasi-peminjaman.html?id=${encodeURIComponent(
              p.kode_peminjaman || ''
            )}" class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
              Review
            </a>
          </div>
        </div>
      `
        )
        .join('');
    } else {
      container.innerHTML = '<p class="text-gray-500">Tidak ada pengajuan pending</p>';
    }
  } catch (error) {
    console.error('Error loading pending:', error);
    const container = document.getElementById('pendingList');
    if (container) container.innerHTML = '<p class="text-red-600">Gagal memuat data</p>';
  }
}



