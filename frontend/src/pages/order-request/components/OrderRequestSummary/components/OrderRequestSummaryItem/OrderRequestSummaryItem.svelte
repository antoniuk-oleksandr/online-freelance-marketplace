<script lang="ts">
    type OrderRequestSummaryItemProps = {
        title: string,
        value: string | number,
        styles?: string,
    };

    const {title, value, styles}: OrderRequestSummaryItemProps = $props();
</script>

<div class="{styles} flex capitalize items-center justify-between text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    <span>{title}</span>
    <span>{value}</span>
</div>