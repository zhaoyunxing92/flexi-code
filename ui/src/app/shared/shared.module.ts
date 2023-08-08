import { NgModule } from '@angular/core';
import { CommonModule, NgForOf, NgIf, NgSwitch, NgSwitchCase } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatInputModule } from '@angular/material/input';
import { ReactiveFormsModule } from '@angular/forms';
import { MatNativeDateModule } from '@angular/material/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MainRoutingModule } from '../main/main-routing.module';
import { ComponentsModule } from '../components/components.module';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatGridListModule } from '@angular/material/grid-list';

@NgModule({
  imports: [
    NgForOf,
    NgSwitch,
    NgSwitchCase,
    NgIf,
    CommonModule,
    BrowserModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatInputModule,
    MatNativeDateModule,
    MatDatepickerModule,
    MainRoutingModule,
    MatCardModule,
    MatButtonModule,
    MatGridListModule,
    MatInputModule
  ],
  exports: [
    NgForOf,
    NgSwitch,
    NgSwitchCase,
    NgIf,
    CommonModule,
    BrowserModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatInputModule,
    MatNativeDateModule,
    MatDatepickerModule,
    MainRoutingModule,
    MatCardModule,
    MatButtonModule,
    MatGridListModule,
    MatInputModule
  ]
})
export class SharedModule {
}
