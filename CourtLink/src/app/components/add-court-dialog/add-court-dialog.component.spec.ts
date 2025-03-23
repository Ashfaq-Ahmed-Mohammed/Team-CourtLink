import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddCourtDialogComponent } from './add-court-dialog.component';

describe('AddCourtDialogComponent', () => {
  let component: AddCourtDialogComponent;
  let fixture: ComponentFixture<AddCourtDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddCourtDialogComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddCourtDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
