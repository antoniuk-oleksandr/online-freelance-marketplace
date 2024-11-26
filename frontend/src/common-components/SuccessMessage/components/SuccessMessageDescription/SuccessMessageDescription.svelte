<script lang="ts">
    type SuccessMessageDescriptionProps = {
        description: string,
    }

    const {description}: SuccessMessageDescriptionProps = $props();
</script>

<p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    {description}
</p>