<script lang="ts">
  import HeroSection from '@/pages/home/components/HomeHeroSection/HomeHeroSection.svelte'
  import HomePageLayout from '@/pages/home/HomePageLayout.svelte'
  import CategoriesSection from '@/pages/home/components/HomeCategoriesSection/HomeCategoriesSection.svelte'
  import StatsSection from '@/pages/home/components/HomeStatsSection/HomeStatsSection.svelte'
  import HowItWorksSection from '@/pages/home/components/HomeHowItWorksSection/HomeHowItWorksSection.svelte'
  import HomeReviewsSection from '@/pages/home/components/HomeReviewsSection/HomeReviewsSection.svelte'
  import HomeBestFreelancers from '@/pages/home/components/HomeBestFreelancersSection/HomeBestFreelancersSection.svelte'
  import HomeFAQSection from '@/pages/home/components/HomeFAQSection/HomeFAQSection.svelte'
  import HomeSecuritySection from '@/pages/home/components/HomeSecuritySection/HomeSecuritySection.svelte'
  import HomeProjectsSection from '@/pages/home/components/HomeProjectsSection/HomeProjectsSection.svelte'
  import { request } from '@/api/request'
  import type { GetHomeDataRequestResponse } from '@/types/GetHomeDataRequestResponse.ts'
  import { errorStore } from '@/common-stores/error-store'
  import type { HomeData } from '@/types/HomeData'

  let homeData = $state<HomeData | undefined>()

  request<GetHomeDataRequestResponse>('GET', '/home-data').then((response) => {
    if (response.status === 200) homeData = response.data
    else {
      console.log(response)
      errorStore.set({ error: response.data.error, shown: true })
    }
  })
</script>

{#if homeData}
  <HomePageLayout>
    <HeroSection />
    <CategoriesSection />
    <HowItWorksSection />
    <HomeBestFreelancers bestFreelancers={homeData.bestFreelancers} />
    <HomeProjectsSection bestServices={homeData.bestServices} />
    <HomeReviewsSection bestReviews={homeData.bestReviews} />
    <HomeSecuritySection />
    <StatsSection keyMetrics={homeData.keyMetrics} />
    <HomeFAQSection />
  </HomePageLayout>
{/if}
