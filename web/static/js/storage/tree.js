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

    nameInput.value = "";
    descriptionInput.value = "";
    render();
    closeFolderModal();
}

async function uploadFile() {

    const nameInput =
        document.getElementById("file-name");

    const descriptionInput =
        document.getElementById("upload-description");

    const name =
        nameInput.value.trim();

    if(!name){
        alert("Enter file name");
        return;
    }

    if(!selectedFile){
        alert("Select file");
        return;
    }

    const form = new FormData();

    form.append(
        "file",
        selectedFile
    );

    form.append(
        "name",
        name
    );

    form.append(
        "description",
        descriptionInput.value.trim()
    );

    form.append(
        "parent_id",
        tree.node.id
    );

    const response = await api(
        "/api/storage/files",
        {
            method:"POST",
            body:form
        }
    );

    if(!response.ok){

        const error =
            await response.text();
        console.error(error);
        alert("Upload failed");
        return;
    }

    const node = await response.json();

    tree.children.push(node);

    nameInput.value="";
    descriptionInput.value="";

    selectedFile=null;

    document.getElementById(
        "upload-file"
    ).value="";

    document.getElementById(
        "filePicker"
    ).textContent="📄 Select File";

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

async function downloadFile(node){

    const response = await api(
        "/api/storage/files/" 
        + node.id 
        + "/download"
    );
    if(!response.ok){

        const error =
            await response.text();

        console.error(error);

        alert("Download failed");

        return;
    }
    const blob =
        await response.blob();
    const url =
        window.URL.createObjectURL(blob);

    const a =
        document.createElement("a");
    a.href = url;
    let filename = node.name;
    if(node.extension){
        filename += node.extension;
    }
    a.download = filename;
    document.body.appendChild(a);
    a.click();
    a.remove();
    window.URL.revokeObjectURL(url);
}

function downloadSelectedNode(){

    if(!selectedNode)
        return;
    downloadFile(selectedNode);
}

