// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { CursorAPI } from './cursor-api'
import { Cursor, CopyCursorToCursorAPI } from './cursor'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class CursorService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CursorServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private cursorsUrl: string

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
    this.cursorsUrl = origin + '/api/github.com/thomaspeugeot/phyllotaxymusic/substackcursor/go/v1/cursors';
  }

  /** GET cursors from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI[]> {
    return this.getCursors(GONG__StackPath, frontRepo)
  }
  getCursors(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<CursorAPI[]>(this.cursorsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<CursorAPI[]>('getCursors', []))
      );
  }

  /** GET cursor by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI> {
    return this.getCursor(id, GONG__StackPath, frontRepo)
  }
  getCursor(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.cursorsUrl}/${id}`;
    return this.http.get<CursorAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched cursor id=${id}`)),
      catchError(this.handleError<CursorAPI>(`getCursor id=${id}`))
    );
  }

  // postFront copy cursor to a version with encoded pointers and post to the back
  postFront(cursor: Cursor, GONG__StackPath: string): Observable<CursorAPI> {
    let cursorAPI = new CursorAPI
    CopyCursorToCursorAPI(cursor, cursorAPI)
    const id = typeof cursorAPI === 'number' ? cursorAPI : cursorAPI.ID
    const url = `${this.cursorsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CursorAPI>(url, cursorAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CursorAPI>('postCursor'))
    );
  }
  
  /** POST: add a new cursor to the server */
  post(cursordb: CursorAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI> {
    return this.postCursor(cursordb, GONG__StackPath, frontRepo)
  }
  postCursor(cursordb: CursorAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CursorAPI>(this.cursorsUrl, cursordb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted cursordb id=${cursordb.ID}`)
      }),
      catchError(this.handleError<CursorAPI>('postCursor'))
    );
  }

  /** DELETE: delete the cursordb from the server */
  delete(cursordb: CursorAPI | number, GONG__StackPath: string): Observable<CursorAPI> {
    return this.deleteCursor(cursordb, GONG__StackPath)
  }
  deleteCursor(cursordb: CursorAPI | number, GONG__StackPath: string): Observable<CursorAPI> {
    const id = typeof cursordb === 'number' ? cursordb : cursordb.ID;
    const url = `${this.cursorsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CursorAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted cursordb id=${id}`)),
      catchError(this.handleError<CursorAPI>('deleteCursor'))
    );
  }

  // updateFront copy cursor to a version with encoded pointers and update to the back
  updateFront(cursor: Cursor, GONG__StackPath: string): Observable<CursorAPI> {
    let cursorAPI = new CursorAPI
    CopyCursorToCursorAPI(cursor, cursorAPI)
    const id = typeof cursorAPI === 'number' ? cursorAPI : cursorAPI.ID
    const url = `${this.cursorsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<CursorAPI>(url, cursorAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CursorAPI>('updateCursor'))
    );
  }

  /** PUT: update the cursordb on the server */
  update(cursordb: CursorAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI> {
    return this.updateCursor(cursordb, GONG__StackPath, frontRepo)
  }
  updateCursor(cursordb: CursorAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CursorAPI> {
    const id = typeof cursordb === 'number' ? cursordb : cursordb.ID;
    const url = `${this.cursorsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CursorAPI>(url, cursordb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated cursordb id=${cursordb.ID}`)
      }),
      catchError(this.handleError<CursorAPI>('updateCursor'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CursorService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CursorService" + error); // log to console instead

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