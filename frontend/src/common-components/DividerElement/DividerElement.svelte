<script lang="ts">
    type DividerElementProps = {
        orientation?: "horizontal" | "vertical",
        className?: string,
    };

    const {orientation, className}: DividerElementProps = $props();

    let styles = $state('');
    if (orientation === "horizontal" || !orientation) styles = 'border-t w-full';
    else styles = 'border-l h-full';
</script>

<div
        class={`${className} ${styles} border-light-palette-divider dark:border-dark-palette-divider`}
></div>