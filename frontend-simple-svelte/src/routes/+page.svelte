<script>
    import {goto} from "$app/navigation";

    /** @type {import('./$types').PageData} */
    export let data;

    function handleRowClick(petName, petId) {
        goto(`/view/${petName}?id=${petId}`);
    }
</script>

<div class="container">
    <h1>My pets</h1>


<div class="table-responsive">
    <table class="table table-striped table-hover">
        <thead>
        <tr>
            <th scope="col">#</th>
            <th scope="col">Name</th>
            <th scope="col">Type</th>
            <th scope="col">Id</th>
            <th scope="col">Profile Pic</th>
        </tr>
        </thead>
        <tbody>
        {#each data.pets as pet, idx}
            <tr on:click={() => handleRowClick(pet.name, pet.id)} >
                <th scope="row">{idx + 1}</th>
                <td>{pet.name}</td>
                <td>{pet.type}</td>
                <td class="vermall">{pet.id}</td>
                <td>
                    <img src={pet.profilePictureUrl}  class="img-thumbnail" style="max-width: 100px; max-height: 100px;">
                </td>
            </tr>
        {/each}
        </tbody>
    </table>
</div>
</div>