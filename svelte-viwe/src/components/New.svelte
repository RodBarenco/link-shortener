<script>
    import { Modals, closeModal, openModal } from "svelte-modals";
    import Modal from "./Modal.svelte";

    async function createGoly(data) {
        const json = {
            redirect: data.redirect,
            goly: data.goly,
            random: data.random
        };
        await fetch("http://localhost:8000/goly", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(json)
        }).then(res => {
            console.log(res);
            closeModal();
            location.reload();
        });
    }

    function handleOpen() {
        const redirect = ""; // Preencha com o valor correto
        const goly = ""; // Preencha com o valor correto
        const random = false; // Preencha com o valor correto

        openModal(Modal, {
            title: "Create New Goaly Link",
            send: createGoly,
            redirect,
            goly,
            random
        });
    }
</script>

<button on:click={handleOpen}>New +</button>

<style>
    button {
        background-color: greenyellow;
        color: white;
        font-size: 28px;
        border: 5px;
        padding: .75rem;
        border-radius: 5px;
    }
</style>
