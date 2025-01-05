<script lang="ts">
    import Icon from "@iconify/svelte";
    import HomeBestFreelancersSectionItemProjectsNumLayout
        from "@/pages/home/components/HomeBestFreelancersSection/components/HomeBestFreelancersSectionItemProjectsNum/HomeBestFreelancersSectionItemProjectsNumLayout.svelte";
    import type {BestFreelancer} from "@/types/BestFreelancer.ts";

    const {completedProjects}: BestFreelancer = $props();
</script>

<HomeBestFreelancersSectionItemProjectsNumLayout>
    <Icon icon="hugeicons:checkmark-circle-01" class="text-cyan-500 dark:text-cyan-500" width="20" height="20"/>
    <span class="text-sm text-light-palette-text-primary dark:text-dark-palette-text-primary">
        {completedProjects} Projects Completed
    </span>
</HomeBestFreelancersSectionItemProjectsNumLayout>

