<script lang="ts">
    import Icon from "@iconify/svelte";
    import CheckboxLayout from "@/common-components/Checkbox/CheckboxLayout.svelte";

    type CheckboxProps = {
        selected: boolean,
    }

    const {selected} : CheckboxProps = $props();
</script>

<CheckboxLayout>
    {#if selected}
        <Icon width={18} height={18} icon="material-symbols:check"/>
    {/if}
</CheckboxLayout>