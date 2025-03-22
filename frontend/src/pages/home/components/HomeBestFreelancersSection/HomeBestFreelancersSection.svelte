<script lang="ts">
  import HomeSectionLayout from '@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte'
  import HomeSectionHeader from '@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte'
  import HomeBestFreelancersSectionItemList from '@/pages/home/components/HomeBestFreelancersSection/components/HomeBestFreelancersSectionItemList/HomeBestFreelancersSectionItemList.svelte'
  import type { BestFreelancer } from '@/types/BestFreelancer.ts'

  type HomeBestFreelancersSectionProps = {
    bestFreelancers: BestFreelancer[]
  }

  const props: HomeBestFreelancersSectionProps = $props()
</script>

<HomeSectionLayout>
  <HomeSectionHeader
    title="Featured Freelancers"
    description="Meet some of the top-rated professionals on our platform."
  />
  <HomeBestFreelancersSectionItemList {...props} />
</HomeSectionLayout>
