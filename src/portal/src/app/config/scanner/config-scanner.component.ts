import {Component, ViewChild, OnInit} from "@angular/core";
import {Scanner, ScannerMetadata} from "./scanner";
import {NewScannerModalComponent} from "./new-scanner-modal.component";
import {ConfigScannerService} from "./config-scanner.service";
import {Observable} from "rxjs";
import {scan} from "rxjs/operators";

@Component({
    selector: 'config-scanner',
    templateUrl: "config-scanner.component.html",
    styleUrls: ['./config-scanner.component.scss', '../config.component.scss']
})
export class ConfigurationScannerComponent implements OnInit {

    defaultScanner: Scanner;
    scannerMetadata: ScannerMetadata;

    scanners: Scanner[] = [];

    selectedRows: Scanner[] = [];

    @ViewChild(NewScannerModalComponent)
    newScannerDialog: NewScannerModalComponent;

    constructor(
        private configScannerService: ConfigScannerService
    ) {
    }

    ngOnInit() {
        this.refreshScanners();
    }

    private refreshScanners(): void {
        this.selectedRows = [];
        this.getDefaultScanner();
        this.getScanners();
    }

    private getDefaultScanner(): void {
        this.configScannerService.getDefault().subscribe(scanner => {
            this.defaultScanner = scanner;
            this.getScannerMetadata(scanner.id);
        }, error => {
            alert('Error while getting default scanner registration: ' + error);
        });
    }

    private getScannerMetadata(id: number): void {
        this.configScannerService.getMetadata(id).subscribe(metadata => {
            this.scannerMetadata = metadata;
        }, error => {
            alert('Error while getting scanner metadata:' + error);
        });
    }

    private getScanners(): void {
        this.configScannerService.getScanners().subscribe(scanners => {
            this.scanners = scanners;
        }, error => {
            alert('Error while getting scanner registrations: ' + error);
        });
    }

    public get isDisabledSetAsDefault(): boolean {
        if (!this.isSingleRowSelected()) {
            return true;
        }
        return this.selectedRows[0].default === true;
    }

    setAsDefault(scanner: Scanner) {
        this.configScannerService.setAsDefault(scanner.id).subscribe(() => {
            this.refreshScanners();
        }, error => {
            alert('Error while setting scanner registration as default: ' + error);
        });
    }

    public get isDisabledEdit(): boolean {
        return !this.isSingleRowSelected();
    }

    edit(selectedScanner: Scanner) {
        this.configScannerService.getScanner(selectedScanner.id).subscribe(scanner => {
            this.newScannerDialog.open(scanner);
        }, error => {
            alert('Error while getting scanner registration: ' + error);
        });
    }

    public get isDisabledDelete(): boolean {
        if (!this.isSingleRowSelected()) {
            return true;
        }
        return this.selectedRows[0].default === true;
    }

    delete(scanner: Scanner) {
        this.configScannerService.delete(scanner.id).subscribe(() => {
            this.refreshScanners();
        }, error => {
            alert("Error while deleting scanner registration: " + error);
        });
    }

    selectedChange(): void {

    }

    newScanner(): void {
        let scanner = new Scanner();
        scanner.enabled = true;
        this.newScannerDialog.open(scanner);
    }

    onScannerSaved(scanner: Scanner): void {
        this.refreshScanners();
    }

    private isSingleRowSelected(): boolean {
        return this.selectedRows.length === 1;
    }

}
