<script lang="ts">
    type HowItWorksSectionStepDescriptionProps = {
        description: string;
    };

    const {description}: HowItWorksSectionStepDescriptionProps = $props();
</script>

<p class="text-sm text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    {description}
</p>