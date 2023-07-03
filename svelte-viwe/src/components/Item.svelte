<script>
    import Card from "./Card.svelte";
    import Modal from "./Modal.svelte";
    import { Modals, closeModal, openModal } from "svelte-modals";
   
    export let goly;
    let ifShow = true;

    async function update(data) {
    if (confirm("Are you sure you wish to update this goly link?")) {
        const json = {
            redirect: data.redirect,
            goly: data.goly,
            random: data.random,
            id: goly.id
        };

        try {
            const response = await Promise.race([
                fetch("http://localhost:8000/goly", {
                    method: "PATCH",
                    headers: { "Content-type": "application/json" },
                    body: JSON.stringify(json)
                }),
                new Promise((_, reject) =>
                    setTimeout(() => reject(new Error("Timeout")), 5000) // Tempo limite de 5 segundos
                )
            ]);

            if (response.ok) {
                const updatedData = await response.json();

                goly.redirect = updatedData.redirect;
                goly.goly = updatedData.goly;
                goly.random = updatedData.random;

                closeModal();

                console.log("Goly link updated successfully!");
                console.log("Updated Data:", updatedData);
            } else {
                console.log("Error updating the goly link");
            }
        } catch (error) {
            console.log("How to handle a communication failure with the database:", error.message);
            }
        }
    }


    function handleOpen(goly) {
        openModal(Modal, {
            title: "Update Goly link",
            send: update,
            goly: goly.goly,
            redirect: goly.redirect,
            random: goly.random
        });
    }

    async function deleteGoly () {
        if (confirm("Are you sure you wish to delete this goly link (" + goly.goly + ")?")) {
            await fetch("http://localhost:8000/goly/" + goly.id, {
                method: "DELETE"
            }).then(res => {
            ifShow = false;
            console.log(res)
            })
        }
    }

</script>
{#if ifShow}
<Card>
    <container class="card-item-container">
        <p>Goly: <a href="http://localhost:8000/voyage/{goly.goly}" 
        target="_blank" style="color: beige;">http://localhost:8000/voyage/{goly.goly}</a></p>
        
        <p>Redirect: {goly.redirect}</p>
        <p>Clicked: {goly.clicked}</p>
        <button class="update" on:click={() => handleOpen(goly)}>Update</button>
        <button class="delete" on:click={() => deleteGoly(goly)}>Delete</button>
    </container>
</Card>

{/if}
<Modals>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div
        slot="backdrop"
        class="backdrop"
        on:click={closeModal}
    />
</Modals>

<style>
    .card-item-container{
        width: 100%;
        height: 100%;
        flex-wrap: wrap;
        word-wrap: break-word;
    }
    
    button {
        color: beige;
        font-weight: bolder;
        border: none;
        padding: .75rem;
        border: 15px;
        border-radius: 5px;
    }

    .update {
        background-color: greenyellow;
    }

    .delete {
        background-color: crimson;
    }

</style>
