<script lang="ts">
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";
    import LeftSignSideOrTextLayout
        from "@/common-components/Sign/components/LeftSignSideOrText/LeftSignSideOrTextLayout.svelte";
    import type {SignProps} from "@/types/SignProps.ts";

    const {subFormText}: SignProps = $props();
</script>

{#if subFormText !== undefined}
    <LeftSignSideOrTextLayout>
        <DividerElement/>
        <div class="absolute -top-3 font-medium text-base left-0 text-center w-full">
            <span class={"bg-light-palette-background-default dark:bg-dark-palette-background-default lg:bg-light-palette-background-block lg:dark:bg-dark-palette-background-block px-2.5"}>or</span>
        </div>
    </LeftSignSideOrTextLayout>
{/if}
