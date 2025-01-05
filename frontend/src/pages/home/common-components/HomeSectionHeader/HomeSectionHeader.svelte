<script lang="ts">
    type HomeSectionHeader = {
        title: string,
        description: string,
    }

    const {title, description}: HomeSectionHeader = $props();
</script>

<div class="flex flex-col gap-y-3">
    <h2 class="text-3xl font-semibold text-cyan-500">{title}</h2>
    <p class="text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
        {description}
    </p>
</div>