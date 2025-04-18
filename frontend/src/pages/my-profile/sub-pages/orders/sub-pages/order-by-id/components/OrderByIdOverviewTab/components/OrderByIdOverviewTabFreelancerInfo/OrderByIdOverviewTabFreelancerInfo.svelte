<script lang="ts">
  import { Link } from 'svelte-routing'

  type OrderByIdOverviewTabFreelancerInfoProps = {
    freelancer: {
      username: string
      id: number
    }
  }

  const { freelancer }: OrderByIdOverviewTabFreelancerInfoProps = $props()
</script>

<div class="flex gap-1 flex-col md:flex-row">
  <span>Ordered from</span>
  <Link to="/users/{freelancer.id}">
    <span class="hover:underline text-cyan-500 font-medium">{freelancer.username}</span>
  </Link>
</div>
