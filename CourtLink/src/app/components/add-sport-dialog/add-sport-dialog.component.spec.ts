import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddSportDialogComponent } from './add-sport-dialog.component';

describe('AddSportDialogComponent', () => {
  let component: AddSportDialogComponent;
  let fixture: ComponentFixture<AddSportDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddSportDialogComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddSportDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
