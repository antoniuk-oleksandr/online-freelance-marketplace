<script lang="ts">
    import Icon from "@iconify/svelte";
    import KeepMeSignedInBlockIconLayout
        from "@/pages/sign-in/components/KeepMeSignedInBlock/components/KeepMeSignedInBlockIcon/KeepMeSignedInBlockIconLayout.svelte";

    type KeepMeSignedInBlockIconProps = {
        keepSignedIn: boolean,
    }

    const {keepSignedIn}: KeepMeSignedInBlockIconProps = $props();
</script>

<KeepMeSignedInBlockIconLayout>
    {#if keepSignedIn}
        <Icon width={18} height={18} icon="material-symbols:check"/>
    {/if}
</KeepMeSignedInBlockIconLayout>
