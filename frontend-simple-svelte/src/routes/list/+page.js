import { PetServiceApiClient } from '$lib/petServiceApiClient.js';
import {ListPetsResponse} from "../../proto/api.ts";


const client = new PetServiceApiClient();

export async function load({ request }) {
    // const authorization = request.headers.get('authorization');
    // const meta = {
    //     authorization: authorization
    // };
    //
    // console.log("Authorization: " + authorization);

    const meta = {};

    try {
        const response = await client.PetServiceClient.listPets({}, { meta });
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
            pets: response.response.pets
        }
    } catch (error) {
        console.error(error);
        return {
            pets: []
        };
    }
}

