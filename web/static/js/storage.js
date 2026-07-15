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
}

async function api(url,options={}) {
    const token =
    sessionStorage.getItem("access_token");
    options.headers = {
        ...(options.headers || {}),
        "Authorization":
        "Bearer " + token
    };
    return fetch(url, {
        credentials:"include",
        ...options
    });
}

async function refresh() {
    const response =
    await fetch("/auth/refresh", {
        method:"POST",
        credentials:"include"
    });
    if (!response.ok) {
        logoutLocal();
        return false;
    }
    const data =
    await response.json();
    sessionStorage.setItem(
        "access_token",
        data.access_token
    );
    return true;
}

async function loadUser() {
    const response =
    await api("/users/me");
    if (!response.ok) {
        logoutLocal();
        return;
    }
    const user =
    await response.json();
    document.querySelector(
        ".user-name"
    ).textContent=user.name;
    document.querySelector(
        ".user-email"
    ).textContent=user.email;
    document.querySelector(
        ".user-created"
    ).textContent=
    new Date(user.created_at)
    .toLocaleDateString();
}

async function logout(){
    await fetch("/auth/logout",{
        method:"POST",
        credentials:"include"
    });
    logoutLocal();
}

function logoutLocal() {
    sessionStorage.removeItem(
        "access_token"
    );
    window.location.href="/";
}

async function initStorage() {
    const ok =
    await refresh();
    if (!ok) return;
    await loadUser();
}

document.addEventListener(
    "DOMContentLoaded",
    initStorage
);
