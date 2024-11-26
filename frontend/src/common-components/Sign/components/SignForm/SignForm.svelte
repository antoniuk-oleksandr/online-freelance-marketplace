<script lang="ts">
    import type {SignProps} from "@/types/SignProps.ts";
    import Button from "@/common-components/Button/Button.svelte";
    import LeftSignSideOrText from "@/common-components/Sign/components/LeftSignSideOrText/LeftSignSideOrText.svelte";
    import SignFormLayout from "@/common-components/Sign/components/SignForm/SignFormLayout.svelte";
    import GoogleButton from "@/common-components/Sign/components/GoogleButton/GoogleButton.svelte";
    import {navigate} from "svelte-routing";
    import BackSignInButton from "@/common-components/Sign/components/BackSignInButton/BackSignInButton.svelte";

    const signProps: SignProps = $props();
    let loading = $state(false);
    const setLoading = (value: boolean) => loading = value;
</script>

<SignFormLayout setLoading={setLoading} {...signProps}>
    {#if signProps.children}
        {@render signProps.children()}
    {/if}
    <Button
            loading={loading}
            styles="text-lg !h-12"
            type={"submit"}
    >{signProps.signButtonText}</Button>
    {#if signProps.showBackButton}
        <BackSignInButton/>
    {/if}
    <LeftSignSideOrText {...signProps}/>
    <GoogleButton {...signProps}/>
</SignFormLayout>
