import { PetServiceApiClient } from '$lib/petServiceApiClient.js';
import {ListMealsRequest} from "$lib/proto/api.ts";

const client = new PetServiceApiClient();

export async function load({ params, url }) {

    const meta = {};
    const listMealsRequest = {
        petId: url.searchParams.get('id')
    };

    try {
        const response = await client.PetServiceClient.getMeals(listMealsRequest ,meta);

        console.log("Response: ", response.response);


        //Convert pets to a serializable format
        // const pets = response.response.pets.map(pet => ({
        //     id: pet.id,
        //     name: pet.name,
        //     type: pet.type,
        //     age: pet.age
        // }));

        // return {
        //     pets: pets
        // };

        return {
            meals: response.response.meals,
            "name": params.slug

        }
    } catch (error) {
        console.error(error);
        return {
            meals: [],
            "name": params.slug
        };
    }
}

