<script lang="ts">
    import type {SignProps} from "@/types/SignProps.ts";
    import Button from "@/common-components/Button/Button.svelte";
    import LeftSignSideOrText from "@/common-components/Sign/components/LeftSignSideOrText/LeftSignSideOrText.svelte";
    import SignFormLayout from "@/common-components/Sign/components/SignForm/SignFormLayout.svelte";
    import GoogleButton from "@/common-components/Sign/components/GoogleButton/GoogleButton.svelte";

    const signProps: SignProps = $props();
</script>

<SignFormLayout {...signProps}>
    {#if signProps.children}
        {@render signProps.children()}
    {/if}
    <Button
            styles="text-lg !h-14"
            type={"submit"}
    >{signProps.signButtonText}</Button>
    <LeftSignSideOrText/>
    <GoogleButton {...signProps}/>
</SignFormLayout>
