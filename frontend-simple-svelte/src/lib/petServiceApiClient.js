import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { PetServiceClient } from "$lib/proto/api.client.ts";

/**
 * The `PetServiceApiClient` class is responsible for creating an instance of the `PetServiceClient`
 * with the appropriate gRPC web transport configuration.
 *
 * This class initializes the gRPC transport with a base URL and provides
 * a `PetServiceClient` instance that can be used to make RPC calls to the pet service.
 */
export class PetServiceApiClient {
    constructor() {
        let transport = new GrpcWebFetchTransport({
            baseUrl: `http://localhost:4180/api/v1`,
        });

        this.PetServiceClient = new PetServiceClient(transport);
    }
}