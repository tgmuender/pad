<script>
    import {goto} from "$app/navigation";

    /** @type {import('./$types').PageData} */
    export let data;

    function handleRowClick(petName, petId) {
        goto(`/view/${petName}?id=${petId}`);
    }
</script>

My pets

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
                <td>{pet.id}</td>
                <td>asdf</td>
            </tr>
        {/each}
        </tbody>
    </table>
</div>