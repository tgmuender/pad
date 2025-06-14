export function load({ params, url }) {


    return {
        "name": params.slug,
        "id": url.searchParams.get('id')
    };
}