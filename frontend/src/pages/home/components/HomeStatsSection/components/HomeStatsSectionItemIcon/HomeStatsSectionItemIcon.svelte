<script lang="ts">
    import Icon from "@iconify/svelte";

    type StatsSectionItemIconProps = {
        icon: string,
    };

    const {icon}: StatsSectionItemIconProps = $props();
</script>

<div class="mt-4">
    <Icon class="text-cyan-500 dark:text-cyan-500" icon={icon} width="48"
          height="48"/>
</div>