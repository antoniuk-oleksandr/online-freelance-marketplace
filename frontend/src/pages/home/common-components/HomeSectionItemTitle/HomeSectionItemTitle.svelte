<script lang="ts">
    type HowItWorksSectionStepTitleProps = {
        title: string;
    };

    const {title}: HowItWorksSectionStepTitleProps = $props();
</script>

<h3 class="text-xl font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary">
    {title}
</h3>
