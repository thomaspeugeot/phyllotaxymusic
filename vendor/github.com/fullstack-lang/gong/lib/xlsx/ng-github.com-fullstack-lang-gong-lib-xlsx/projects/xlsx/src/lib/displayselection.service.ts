// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { DisplaySelectionAPI } from './displayselection-api'
import { DisplaySelection, CopyDisplaySelectionToDisplaySelectionAPI } from './displayselection'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { XLFileAPI } from './xlfile-api'
import { XLSheetAPI } from './xlsheet-api'

@Injectable({
  providedIn: 'root'
})
export class DisplaySelectionService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  DisplaySelectionServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private displayselectionsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.displayselectionsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/xlsx/go/v1/displayselections';
  }

  /** GET displayselections from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI[]> {
    return this.getDisplaySelections(Name, frontRepo)
  }
  getDisplaySelections(Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<DisplaySelectionAPI[]>(this.displayselectionsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<DisplaySelectionAPI[]>('getDisplaySelections', []))
      );
  }

  /** GET displayselection by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI> {
    return this.getDisplaySelection(id, Name, frontRepo)
  }
  getDisplaySelection(id: number, Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.displayselectionsUrl}/${id}`;
    return this.http.get<DisplaySelectionAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched displayselection id=${id}`)),
      catchError(this.handleError<DisplaySelectionAPI>(`getDisplaySelection id=${id}`))
    );
  }

  // postFront copy displayselection to a version with encoded pointers and post to the back
  postFront(displayselection: DisplaySelection, Name: string): Observable<DisplaySelectionAPI> {
    let displayselectionAPI = new DisplaySelectionAPI
    CopyDisplaySelectionToDisplaySelectionAPI(displayselection, displayselectionAPI)
    const id = typeof displayselectionAPI === 'number' ? displayselectionAPI : displayselectionAPI.ID
    const url = `${this.displayselectionsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<DisplaySelectionAPI>(url, displayselectionAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<DisplaySelectionAPI>('postDisplaySelection'))
    );
  }
  
  /** POST: add a new displayselection to the server */
  post(displayselectiondb: DisplaySelectionAPI, Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI> {
    return this.postDisplaySelection(displayselectiondb, Name, frontRepo)
  }
  postDisplaySelection(displayselectiondb: DisplaySelectionAPI, Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<DisplaySelectionAPI>(this.displayselectionsUrl, displayselectiondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted displayselectiondb id=${displayselectiondb.ID}`)
      }),
      catchError(this.handleError<DisplaySelectionAPI>('postDisplaySelection'))
    );
  }

  /** DELETE: delete the displayselectiondb from the server */
  delete(displayselectiondb: DisplaySelectionAPI | number, Name: string): Observable<DisplaySelectionAPI> {
    return this.deleteDisplaySelection(displayselectiondb, Name)
  }
  deleteDisplaySelection(displayselectiondb: DisplaySelectionAPI | number, Name: string): Observable<DisplaySelectionAPI> {
    const id = typeof displayselectiondb === 'number' ? displayselectiondb : displayselectiondb.ID;
    const url = `${this.displayselectionsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<DisplaySelectionAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted displayselectiondb id=${id}`)),
      catchError(this.handleError<DisplaySelectionAPI>('deleteDisplaySelection'))
    );
  }

  // updateFront copy displayselection to a version with encoded pointers and update to the back
  updateFront(displayselection: DisplaySelection, Name: string): Observable<DisplaySelectionAPI> {
    let displayselectionAPI = new DisplaySelectionAPI
    CopyDisplaySelectionToDisplaySelectionAPI(displayselection, displayselectionAPI)
    const id = typeof displayselectionAPI === 'number' ? displayselectionAPI : displayselectionAPI.ID
    const url = `${this.displayselectionsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<DisplaySelectionAPI>(url, displayselectionAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<DisplaySelectionAPI>('updateDisplaySelection'))
    );
  }

  /** PUT: update the displayselectiondb on the server */
  update(displayselectiondb: DisplaySelectionAPI, Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI> {
    return this.updateDisplaySelection(displayselectiondb, Name, frontRepo)
  }
  updateDisplaySelection(displayselectiondb: DisplaySelectionAPI, Name: string, frontRepo: FrontRepo): Observable<DisplaySelectionAPI> {
    const id = typeof displayselectiondb === 'number' ? displayselectiondb : displayselectiondb.ID;
    const url = `${this.displayselectionsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<DisplaySelectionAPI>(url, displayselectiondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated displayselectiondb id=${displayselectiondb.ID}`)
      }),
      catchError(this.handleError<DisplaySelectionAPI>('updateDisplaySelection'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in DisplaySelectionService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("DisplaySelectionService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
