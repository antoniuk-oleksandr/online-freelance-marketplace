<script lang="ts">
    type DividerElementProps = {
        orientation?: "horizontal" | "vertical",
        styles?: string,
    };

    const {orientation, styles}: DividerElementProps = $props();

    let classStyles = $state('');
    if (orientation === "horizontal" || !orientation) classStyles = 'border-t w-full';
    else classStyles = 'border-l h-full';
</script>

<div
        class={`${styles} ${classStyles} border-light-palette-divider dark:border-dark-palette-divider`}
></div>