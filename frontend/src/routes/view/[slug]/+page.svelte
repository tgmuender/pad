<script>
    import { PetServiceApiClient } from '$lib/petServiceApiClient.js';
    import {SetProfilePictureRequest} from "$lib/proto/api.ts";

    const client = new PetServiceApiClient();

    export let data;

    let file;

    function handleFileChange(event) {
        file = event.target.files[0];
        console.log("Selected file:", file);

        const spp = SetProfilePictureRequest.create({petId: data.id, filename: file.name});
        client.PetServiceClient.setProfilePicture(spp)
            .then(response => {
                console.log("Profile picture set successfully:", response);

                const url = new URL(response.response.uploadUrl);


                fetch(url.toString(), {
                    method: 'PUT',
                    body: file,
                    headers: {
                        'Content-Type': file.type
                    }
                })
                    .then(res => {
                        if (!res.ok) throw new Error('Upload failed');
                        // handle success
                    })
                    .catch(err => {
                        // handle error
                    });
            })
            .catch(error => {
                console.error("Error setting profile picture:", error);
            });
        // You can now upload `file` using fetch or another method
    }
</script>

<div class="container">
    <h1>{data.name}</h1>
    <small>{data.id}</small>


    <div class="card" style="width: 18rem;">

        <div class="card-body">
            <h5 class="card-title">Meals</h5>
            <p class="card-text">View and create detailed meal plans for {data.name}.</p>
            <a href="{data.name}/meals?id={data.id}" class="btn btn-primary">Meals</a>
        </div>
    </div>

    <div class="card" style="width: 18rem;">

        <div class="card-body">
            <h5 class="card-title">Profile Pic</h5>
            <input type = "file" on:change={handleFileChange} />
            {#if file}
                <p>Selected file: {file.name}</p>
            {/if}
        </div>
    </div>

</div>
