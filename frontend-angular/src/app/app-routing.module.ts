import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RegistryComponent } from "./registry/registry.component";
import { BrowseComponent } from "./browse/browse.component";
import { SearchComponent } from "./search/search.component";


const routes: Routes = [
  { path: "",redirectTo:"registry", pathMatch: "full" },
  { path: "registry", component: RegistryComponent },
  { path: "browse", component: BrowseComponent },
  { path: "search", component: SearchComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
