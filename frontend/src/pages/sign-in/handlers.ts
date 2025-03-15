import { formStore } from "@/common-components/Sign/stores/form-store";

export const handleSignIn = async (data: any) => {
    console.log(data);
}

export const handleKeepsSignedInClick = (
    keepSignedIn: boolean,
    setKeepSignedIn: (value: boolean) => void
) => {
    formStore.update((prev) => ({ ...prev, keepSignedIn: !keepSignedIn }))
    setKeepSignedIn(!keepSignedIn);
}
