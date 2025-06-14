<script>
    import { PetServiceApiClient } from '$lib/petServiceApiClient.js';
    import {NewPetRequest, Sex} from '$lib/proto/api.ts';
    import { Timestamp } from '$lib/proto/google/protobuf/timestamp.ts';

    const client = new PetServiceApiClient();

    let name = '';
    let type = '';
    let description = '';
    let nameError = '';
    let typeError = '';
    let submitError = '';
    let dob = '';
    let sex = '';

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

    function asTimestamp(dateString) {
        if (!dateString) return null;
        if (typeof dateString !== 'string') {
            console.error('Invalid date string:', dateString);
            return null;
        }
        const date = new Date(dateString);
        if (isNaN(date.getTime())) {
            console.error('Invalid date:', dateString);
            return null;
        }
        return Timestamp.create({
            seconds: Math.floor(date.getTime() / 1000)
        });
    }

    function asSex(value) {
        if (!value || typeof value !== 'string') {
            console.error('Invalid sex value:', value);
            return Sex.UNKNOWN;
        }
        switch (value) {
            case "Female": return Sex.FEMALE
            case "Male": return Sex.MALE;
            default: return Sex.UNKNOWN;
        }
    }

    async function handleSubmit(event) {
        event.preventDefault();
        if (validateForm()) {
            console.log('Form submitted:', {name, type});

            try {
                const npr = NewPetRequest.create({name, type, description: description, dob: asTimestamp(dob), sex: asSex(sex)});
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
                <br/>

                <label for="description">Please provide a short description of your pet.</label>
                <input type="text" class="form-control" id="description" bind:value={description} placeholder="e.g. playful sofa king" >
                <br/>

                <label for="dob">Date of birth</label>
                <input type="date" class="form-control" id="dob" bind:value={dob} placeholder="YYYY-MM-DD">
                <br/>

                <label for="sex">Sex</label>
                <select id="sex" class="form-select" bind:value={sex}>
                    <option value="Female">Female</option>
                    <option value="Male">Male</option>
                    <option value="Unknown">Unknown</option>
                </select>
                <br/>

            </div>
            <button type="submit" class="btn btn-primary">Finish Setup</button>


            <div class="invalid-feedback d-block">
                {submitError}
            </div>


        </form>
    </div>
</div>
