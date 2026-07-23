function toggleMenu() {
    document
        .getElementById("accountMenu")
        .classList
        .toggle("show");
}

window.onclick = function(e) {
    if (!e.target.closest(".account")) {
        document
            .getElementById("accountMenu")
            .classList
            .remove("show");
    }
};

async function loadUser() {

    const response = await api("/api/users/me");

    if (!response.ok) {
        return;
    }

    const user = await response.json();

    document.querySelector(".user-name").textContent =
        user.name;

    document.querySelector(".user-email").textContent =
        user.email;

    document.querySelector(".user-created").textContent =
        new Date(user.created_at)
            .toLocaleDateString();
}

async function initStorage() {

    const ok = await refresh();

    if (!ok) {
        return;
    }


    await loadUser();
    await loadRootTree();
    render();
}

document.addEventListener(
    "DOMContentLoaded",
    initStorage
);

const uploadInput = document.getElementById("upload-file");
const filePicker = document.getElementById("filePicker");

uploadInput.addEventListener("change", () => {
    if (uploadInput.files.length === 0) {
        filePicker.textContent = "📄 Select File";
        selectedFile = null;
        return;
    }
    selectedFile = uploadInput.files[0];
    filePicker.textContent = "📄 " + uploadInput.files[0].name;
});
