<script lang="ts">

    import Icon from "@iconify/svelte";
    import HomeBestFreelancersSectionItemNameLayout
        from "@/pages/home/components/HomeBestFreelancersSection/components/HomeBestFreelancersSectionItemName/HomeBestFreelancersSectionItemNameLayout.svelte";

    type HomeBestFreelancersSectionItemNameProps = {
        firstName: string,
        surname: string,
        rating: number,
    };

    const {firstName, surname, rating}: HomeBestFreelancersSectionItemNameProps = $props();
</script>

<HomeBestFreelancersSectionItemNameLayout>
    <h3 class="text-xl font-semibold text-cyan-500 dark:text-cyan-500">{firstName} {surname}</h3>
    <Icon icon="mingcute:star-fill" class="text-yellow-500 dark:text-yellow-500" width="20" height="20"/>
    <span class="text-sm text-light-palette-text-primary dark:text-dark-palette-text-primary">{rating}</span>
</HomeBestFreelancersSectionItemNameLayout>