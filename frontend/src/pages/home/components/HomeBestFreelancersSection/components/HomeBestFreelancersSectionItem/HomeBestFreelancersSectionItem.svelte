<script lang="ts">
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import HomeBestFreelancersSectionItemName from '@/pages/home/components/HomeBestFreelancersSection/components/HomeBestFreelancersSectionItemName/HomeBestFreelancersSectionItemName.svelte'
  import HomeBestFreelancersSectionItemProjectsNum from '@/pages/home/components/HomeBestFreelancersSection/components/HomeBestFreelancersSectionItemProjectsNum/HomeBestFreelancersSectionItemProjectsNum.svelte'
  import Avatar from '@/common-components/Avatar/Avatar.svelte'
  import { getHoverClass } from '@/utils/utils'
  import { Link } from 'svelte-routing'
  import type { BestFreelancer } from '@/types/BestFreelancer.ts'

  const props: BestFreelancer = $props()

  const hoverClass = getHoverClass()
</script>

<Link to="/users/{props.id}">
  <PaperElement styles="{hoverClass} flex flex-col items-center size-full justify-center gap-2">
    <Avatar size="24" alt="{props.firstName} ${props.surname}" src={props.avatar} />
    <HomeBestFreelancersSectionItemName {...props} />
    <HomeBestFreelancersSectionItemProjectsNum {...props} />
  </PaperElement>
</Link>
