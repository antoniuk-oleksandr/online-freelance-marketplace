<script lang="ts">
    import Icon from "@iconify/svelte";

    type RadioButtonProps = {
        value: boolean,
        setValue: (newValue: boolean) => void,
        width?: number,
        styles?: string,
        height?: number,
    }

    const {value, setValue, styles, height = 18, width = 18}: RadioButtonProps = $props();
</script>

<button
        class="{styles} text-cyan-500"
        type="button"
        onclick={() => setValue(!value)}
>
    {#if value}
        <Icon width={width} height={height} icon="akar-icons:radio-fill"/>
    {:else}
        <Icon width={width} height={height} icon="akar-icons:radio"/>
    {/if}
</button>