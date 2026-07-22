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


async function api(url, options = {}) {

    const request = () => fetch(url, {
        ...options,
        credentials: "include",
        headers: {
            ...(options.headers || {}),
            Authorization:
                "Bearer " + sessionStorage.getItem("access_token")
        }
    });


    let response = await request();


    if (response.status === 401) {

        const ok = await refresh();

        if (!ok) {
            return response;
        }

        response = await request();
    }


    return response;
}


async function refresh() {

    const response = await fetch("/api/auth/refresh", {
        method: "POST",
        credentials: "include"
    });


    if (!response.ok) {
        logoutLocal();
        return false;
    }


    const data = await response.json();

    sessionStorage.setItem(
        "access_token",
        data.access_token
    );


    return true;
}


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



async function logout() {

    await fetch("/api/auth/logout", {
        method: "POST",
        credentials: "include"
    });


    logoutLocal();
}



async function deleteAccount() {

    const response = await api("/api/users/me", {
        method: "DELETE"
    });


    if (!response.ok) {

        const error = await response.text();

        console.error(error);

        alert("Failed to delete account");

        return;
    }


    logoutLocal();
}



function logoutLocal() {

    sessionStorage.removeItem(
        "access_token"
    );

    window.location.href = "/";
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

function openFolderModal() {

    document
        .getElementById("folderModal")
        .classList
        .add("show");
}

function closeFolderModal() {

    document
        .getElementById("folderModal")
        .classList
        .remove("show");
}

function openUploadModal() {

    document
        .getElementById("uploadModal")
        .classList
        .add("show");
}

function closeUploadModal() {

    document
        .getElementById("uploadModal")
        .classList
        .remove("show");
}

let selectedFile = null;

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

/*  TREE DATA   */
/* REQUEST TREE */

const tmpTree = {
    node: {
        id:             null,
        parent_id:      null,
        name:           "root",
        file_type:      "folder",
        mime_type:      null,
        description:    "root dir",
        size:           0,
        upload_at:      "2026-07-18T08:30:00Z",
        updated_at:     "2026-07-19T012:30:00Z"
    },
    
    children: [
        {
            id:             "f4d79f61-bd4d-4c6c-ae72-6d0d0b2d8a02",
            parent_id:      "f4d79f61-bd4d-4c6c-ae72-6d0d0b2d8a01",
            name:           "Doc",
            file_type:      "folder",
            extension:      null,
            mime_type:      null,
            description:    "documents",
            size:           0,
            upload_at:      "2026-07-2T08:30:00Z",
            updated_at:     "2026-07-12T012:30:00Z"
        },
        {
            id:             "f4d79f61-bd4d-4c6c-ae72-6d0d0b2d8a03",
            parent_id:      "f4d79f61-bd4d-4c6c-ae72-6d0d0b2d8a01",
            name:           "Music",
            file_type:      "file",
            extension:      ".wav",
            mime_type:      "audio/wav",
            description:    "audio",
            size:           322,
            upload_at:      "2026-07-3T08:30:00Z",
            updated_at:     "2026-07-12T012:30:00Z"
        },
    ]
}

let tree = null;
let navigation = [];
let selectedNode = null;

async function loadRootTree() {

    const response = await api("/api/storage/tree");

    if (!response.ok) {
        const error = await response.text();
        console.error(error);
        alert("Failed to load storage tree");
        return;
    }

    tree = await response.json();
    navigation = [tree.node]; 
}

/*   RENDER TREE   */
/* FRONT ITEM DATA */

function renderPath() {
    const path = navigation
        .map(node => node.name)
        .join("/");

    document.querySelector(".path").textContent =
        "Path: " + path + "/";
}

function renderTree() {
    const files = document.querySelector(".files");
     files.innerHTML = "";
    
    tree.children.forEach(node => {
        const item = document.createElement("div");

    const isFolder = node.file_type === "folder";

    item.className =
        `item ${isFolder ? "folder" : "file"}`;


    item.innerHTML = `
        ${isFolder ? "📁" : "📄"}
        <span>${node.name}</span>
    `;


    item.addEventListener("contextmenu", (e)=>{
        e.preventDefault();
        openNodeInfoModal(node);
    });

    item.addEventListener("click", ()=>{
        if(node.file_type === "folder"){
            openFolder(node);
        } else {
            downloadFile(node);
        }
    });



    // довге натискання телефон
    let pressTimer;


    item.addEventListener(
        "touchstart",
        ()=>{

            pressTimer = setTimeout(()=>{

                openNodeInfoModal(node);

            },600);

        }
    );


    item.addEventListener(
        "touchend",
        ()=>{

            clearTimeout(pressTimer);

        }
    );
    files.appendChild(item);        
    });
}

function render() {

    if (!tree) {
        return;
    }

    renderPath();
    renderTree();
}

/* CREATE ITEM */
/*  SETUP ITEM */

async function createFolder() {
    const nameInput = document.getElementById("folder-name");
    const descriptionInput = document.getElementById("folder-description");

    const name = nameInput.value.trim();

    if (!name) {
        alert("Enter folder name");
        return;
    }

    const description = descriptionInput.value.trim();

    const request = {
        name: name,
        description: description,
        parent_id: tree.node.id
    };

    const response = await api("/api/storage/folders", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(request)
    });

    if (!response.ok) {
        const error = await response.text();
        console.error(error);
        alert("Failed to create folder");
        return;
    }

    const node = await response.json();

    tree.children.push(node);
//await openFolder(node);

    nameInput.value = "";
    descriptionInput.value = "";

    render();
    closeFolderModal();
}

function uploadFile() {
    const nameInput = document.getElementById("file-name");
    const descriptionInput = document.getElementById("upload-description");

    let name = nameInput.value.trim();

    if (!name) {
        alert("Enter file name");
        return;
    }

    if (!selectedFile) {
        alert("Select file");
        return;
    } 

    const description = descriptionInput.value.trim();

    const extension = selectedFile.name.substring(
        selectedFile.name.lastIndexOf(".")
    ).toLowerCase();
  
    //name = name.replace(/\.[^/.]+$/, "");

    //const finalName = name + extension;

    const now = new Date().toISOString();

    tree.children.push({
        id: crypto.randomUUID(),
        parent_id: tree.node.id,
        name: name,
        file_type: "file",
        extension: extension,
        mime_type: selectedFile.type,
        description: description,
        size: selectedFile.size,
        upload_at: now,
        updated_at: now
    });

    nameInput.value = "";
    descriptionInput.value = "";

    selectedFile = null;
    document.getElementById("upload-file").value = "";
    document.getElementById("filePicker").textContent = "📄 Select File";

    render();
    closeUploadModal();
}

async function openFolder(node){
    
    const response = await api(
        "/api/storage/tree/" + node.id
    );

    if (!response.ok) {
        alert("Failed to open folder");
        return;
    }

    tree = await response.json();

    navigation.push(tree.node);

    render();
}

async function goBack() {

    if (navigation.length <= 1) {
        return;
    }

    navigation.pop();

    const parent = navigation[navigation.length - 1];

    if (parent.parent_id === null) {
        await loadRootTree();
    } else {

        const response = await api(
            "/api/storage/tree/" + parent.id
        );

        if (!response.ok) {
            alert("Failed to load folder");
            return;
        }

        tree = await response.json();
    }

    render();
}

function downloadFile(node){

    console.log("Download:", node.name);

    // потім:
    // window.location.href =
    // "/api/storage/download/" + node.id;

}

function openNodeInfoModal(node) {

 selectedNode = node;
    const downloadBtn = document.getElementById("node-download-btn");

    
    if(node.file_type === "folder"){

        downloadBtn.style.display="none";

    }
    else{
        downloadBtn.style.display="block";
    
    }

    document.getElementById("node-info-title").textContent =
        node.file_type === "folder"
            ? "📁 Folder Info"
            : "📄 File Info";


    document.getElementById("node-info-name").textContent =
        node.name;


    document.getElementById("node-info-type").textContent =
        node.file_type;


    document.getElementById("node-info-description").textContent =
        node.description || "-";


    document.getElementById("node-info-size").textContent =
        node.size + " bytes";


    document.getElementById("node-info-mime").textContent =
        node.mime_type || "-";


    document.getElementById("node-info-extension").textContent =
        node.extension || "-";


    document.getElementById("node-info-created").textContent =
        new Date(node.upload_at)
            .toLocaleString();


    document
        .getElementById("nodeInfoModal")
        .classList
        .add("show");
}

function downloadSelectedNode(){

    if(!selectedNode)
        return;


    downloadFile(selectedNode);

}

function deleteSelectedNode(){

    if(!selectedNode)
        return;


    const index = tree.children.findIndex(
        n => n.id === selectedNode.id
    );


    if(index !== -1){

        tree.children.splice(index,1);

    }


    closeNodeInfoModal();

    render();

}

function closeNodeInfoModal() {

    document
        .getElementById("nodeInfoModal")
        .classList
        .remove("show");
    selectedNode = null;
}
