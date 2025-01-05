<script lang="ts">
    type StatsSectionItemCenterTextProps = {
        centerText: string,
    };

    const {centerText}: StatsSectionItemCenterTextProps = $props();
</script>

<div class="text-xl text-light-palette-text-secondary dark:text-dark-palette-text-secondary mt-2">
    {centerText}
</div>