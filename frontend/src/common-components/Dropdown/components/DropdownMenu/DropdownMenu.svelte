<script lang="ts">
    import DropdownMenuLayout from "@/common-components/Dropdown/components/DropdownMenu/DropdownMenuLayout.svelte";
    import type {DropdownMenuProps} from "@/types/DropdownMenuProps.ts";
    import DropdownMenuContent
        from "@/common-components/Dropdown/components/DropdownMenuContent/DropdownMenuContent.svelte";

    const props: DropdownMenuProps = $props();
</script>

{#if props.showExitAnimation}
    <DropdownMenuLayout {...props}>
        <DropdownMenuContent {...props}/>
    </DropdownMenuLayout>
{/if}
