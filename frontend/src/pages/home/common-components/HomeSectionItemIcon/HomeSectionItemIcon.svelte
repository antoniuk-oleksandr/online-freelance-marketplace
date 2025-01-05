<script lang="ts">
    import Icon from "@iconify/svelte";

    type HowItWorkSectionStepIconProps = {
        icon: string,
    };

    const {icon}: HowItWorkSectionStepIconProps = $props();
</script>

<Icon
        className="mb-2"
        icon={icon} width="48" height="48" class="text-cyan-500 dark:text-cyan-500"
/>
