import { PetServiceApiClient } from '$lib/petServiceApiClient.js';
import {ListPetsResponse} from "../../proto/api.ts";


const client = new PetServiceApiClient();

export async function load({ request }) {

    const meta = {};

    try {
        const response = await client.PetServiceClient.listPets({}, { meta });
        console.log("Response: ", response.response);

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

