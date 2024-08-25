<script>
    import {goto} from "$app/navigation";

    /** @type {import('./$types').PageData} */
    export let data;

    function handleRowClick(petName) {
        goto(`/view/${petName}`);
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
        </tr>
        </thead>
        <tbody>
        {#each data.pets as pet, idx}
            <tr on:click={() => handleRowClick(pet.name)} >
                <th scope="row">{idx + 1}</th>
                <td>{pet.name}</td>
                <td>{pet.type}</td>
            </tr>
        {/each}
        </tbody>
    </table>
</div>