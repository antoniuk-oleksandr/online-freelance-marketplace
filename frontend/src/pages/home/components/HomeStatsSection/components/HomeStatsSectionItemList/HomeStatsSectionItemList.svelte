<script lang="ts">
    import StatsSectionItemListLayout
        from "@/pages/home/components/HomeStatsSection/components/HomeStatsSectionItemList/HomeStatsSectionItemListLayout.svelte";
    import StatsSectionItem
        from "@/pages/home/components/HomeStatsSection/components/HomeStatsSectionItem/HomeStatsSectionItem.svelte";
    import type {KeyMetrics} from "@/types/KeyMetrics.ts";
    import {formatText} from "@/utils/utils.ts";

    type StatsSectionItemListProps = {
        keyMetrics: KeyMetrics,
    };

    const {keyMetrics}: StatsSectionItemListProps = $props();

    type StatsSectionItemType = {
        topText: string,
        centerText: string,
        icon: string,
    };


    const items: StatsSectionItemType[] = [
        {
            topText: formatText(keyMetrics.freelancesAvailable.toString()),
            centerText: "Freelances Available",
            icon: "hugeicons:user-group"
        },
        {
            topText: formatText(keyMetrics.projectsCompleted.toString()),
            centerText: "Projects Completed",
            icon: "hugeicons:checkmark-circle-01"
        },
        {
            topText: `${keyMetrics.avgRating}/5`,
            centerText: "Average Client Rating",
            icon: "hugeicons:star"
        }
    ]
</script>

<StatsSectionItemListLayout>
    {#each items as item}
        <StatsSectionItem {...item}/>
    {/each}
</StatsSectionItemListLayout>