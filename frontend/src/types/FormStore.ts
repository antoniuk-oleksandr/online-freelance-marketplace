export type FormStore = {
    data: any,
    errors: null | Record<string, any>,
    wasSubmitted: boolean,
    keepSignedIn: boolean,
}