<script lang="ts">
  import CreateService from '../../components/CreateService/CreateService.svelte'
  import CreateServiceMediaFormContent from './components/CreateServiceMediaFormContent/CreateServiceMediaFormContent.svelte'
  import { handleCreateServiceMediaFormSubmit } from './hadlers'
  import { createServicePackagesFormSchema } from './helpers'
</script>

<CreateService
  schema={createServicePackagesFormSchema}
  onSubmit={handleCreateServiceMediaFormSubmit}
  activeStepIndex={3}
>
  <CreateServiceMediaFormContent />
</CreateService>
