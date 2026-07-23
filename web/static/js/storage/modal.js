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

function closeNodeInfoModal() {

    document
        .getElementById("nodeInfoModal")
        .classList
        .remove("show");
    selectedNode = null;
}

async function deleteSelectedNode() {

    if(!selectedNode)
        return;

    const response = await api(
        "/api/storage/nodes/" + selectedNode.id,
        {
            method: "DELETE"
        }
    );

    if (!response.ok) {
        const error = await response.text();
        console.error(error);
        alert("Failed to delete node");
        return;
    }

    const index = tree.children.findIndex(
        n => n.id === selectedNode.id
    );

    if(index !== -1){
        tree.children.splice(index, 1);
    }

    closeNodeInfoModal();
    render();
}

