<script lang="ts">
  import CreateService from '../../components/CreateService/CreateService.svelte'
  import CreateServicePackagesFormContent from './components/CreateServicePackagesFormContent/CreateServicePackagesFormContent.svelte'
  import { handleCreateServicePackagesFormSubmit } from './handlers'
  import { createServicePackagesFormSchema } from './helpers'
</script>

<CreateService
  schema={createServicePackagesFormSchema}
  onSubmit={handleCreateServicePackagesFormSubmit}
  activeStepIndex={1}
>
  <CreateServicePackagesFormContent />
</CreateService>
