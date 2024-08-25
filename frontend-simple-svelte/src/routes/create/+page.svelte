<script>
    import { PetServiceApiClient } from '$lib/petServiceApiClient.js';
    import { NewPetRequest } from '$lib/proto/api.ts';

    const client = new PetServiceApiClient();

    let name = '';
    let type = '';
    let nameError = '';
    let typeError = '';
    let submitError = '';

    function validateName() {
        console.log('validateName');
        nameError = '';
        if (name.trim().length === 0) {
            nameError = 'Please choose a name.';
            return false;
        }
        return true;
    }

    function validateType() {
        console.log('validateType');
        typeError = '';
        if (type.trim().length === 0) {
            typeError = 'Please choose a type.';
            return false
        }
        return true;
    }

    function validateForm() {
        return validateName() && validateType();
    }

    async function handleSubmit(event) {
        event.preventDefault();
        if (validateForm()) {
            console.log('Form submitted:', {name, type});

            try {
                const npr = NewPetRequest.create({name, type});
                client.PetServiceClient.newPet(npr);
                await client.PetServiceClient.newPet(npr);

                window.location.href = '/';
            } catch(error) {
                submitError = 'Apologies, your pet could not created, please try again to continue.';
                console.log(error);
            }
        }
    }

</script>

<div class="container">
    <h1>Create a new pet</h1>
    <br/>
    <div>
        <p>Please enter a few basic information about your pet to get started.</p>
        <br/>
        <form on:submit={handleSubmit} class="row g-3">
            <div>
                <label for="name">What is the name of your pet?</label>
                <input type="text" class="form-control" id="name" bind:value={name} on:blur={validateName}>
                <div class="invalid-feedback d-block">{nameError}</div>
                <br/>

                <label for="type">What type of pet do you want to create?</label>
                <select id="type" class="form-select" bind:value={type} on:blur={validateType}>
                    <option value="Dog">Dog</option>
                    <option value="Cat">Cat</option>
                    <option value="Other">Other</option>
                </select>
                <div class="invalid-feedback d-block">{typeError}</div>

            </div>
            <button type="submit" class="btn btn-primary">Finish Setup</button>


            <div class="invalid-feedback d-block">
                {submitError}
            </div>


        </form>
    </div>
</div>
