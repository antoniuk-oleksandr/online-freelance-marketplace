<script lang="ts">
  import HomeBestFreelancersSectionItemListLayout from '@/pages/home/components/HomeBestFreelancersSection/components/HomeBestFreelancersSectionItemList/HomeBestFreelancersSectionItemListLayout.svelte'
  import HomeBestFreelancersSectionItem from '@/pages/home/components/HomeBestFreelancersSection/components/HomeBestFreelancersSectionItem/HomeBestFreelancersSectionItem.svelte'
  import type { BestFreelancer } from '@/types/BestFreelancer.ts'

  type HomeBestFreelancersSectionItemListProps = {
    bestFreelancers: BestFreelancer[]
  }

  const { bestFreelancers }: HomeBestFreelancersSectionItemListProps = $props()
</script>

<HomeBestFreelancersSectionItemListLayout>
  {#each bestFreelancers as freelancer}
    <HomeBestFreelancersSectionItem {...freelancer} />
  {/each}
</HomeBestFreelancersSectionItemListLayout>
