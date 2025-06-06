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

import { DummyAgentAPI } from './dummyagent-api'
import { DummyAgent, CopyDummyAgentToDummyAgentAPI } from './dummyagent'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class DummyAgentService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  DummyAgentServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private dummyagentsUrl: string

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
    this.dummyagentsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/sim/go/v1/dummyagents';
  }

  /** GET dummyagents from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI[]> {
    return this.getDummyAgents(Name, frontRepo)
  }
  getDummyAgents(Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<DummyAgentAPI[]>(this.dummyagentsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<DummyAgentAPI[]>('getDummyAgents', []))
      );
  }

  /** GET dummyagent by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI> {
    return this.getDummyAgent(id, Name, frontRepo)
  }
  getDummyAgent(id: number, Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.dummyagentsUrl}/${id}`;
    return this.http.get<DummyAgentAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched dummyagent id=${id}`)),
      catchError(this.handleError<DummyAgentAPI>(`getDummyAgent id=${id}`))
    );
  }

  // postFront copy dummyagent to a version with encoded pointers and post to the back
  postFront(dummyagent: DummyAgent, Name: string): Observable<DummyAgentAPI> {
    let dummyagentAPI = new DummyAgentAPI
    CopyDummyAgentToDummyAgentAPI(dummyagent, dummyagentAPI)
    const id = typeof dummyagentAPI === 'number' ? dummyagentAPI : dummyagentAPI.ID
    const url = `${this.dummyagentsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<DummyAgentAPI>(url, dummyagentAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<DummyAgentAPI>('postDummyAgent'))
    );
  }
  
  /** POST: add a new dummyagent to the server */
  post(dummyagentdb: DummyAgentAPI, Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI> {
    return this.postDummyAgent(dummyagentdb, Name, frontRepo)
  }
  postDummyAgent(dummyagentdb: DummyAgentAPI, Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<DummyAgentAPI>(this.dummyagentsUrl, dummyagentdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted dummyagentdb id=${dummyagentdb.ID}`)
      }),
      catchError(this.handleError<DummyAgentAPI>('postDummyAgent'))
    );
  }

  /** DELETE: delete the dummyagentdb from the server */
  delete(dummyagentdb: DummyAgentAPI | number, Name: string): Observable<DummyAgentAPI> {
    return this.deleteDummyAgent(dummyagentdb, Name)
  }
  deleteDummyAgent(dummyagentdb: DummyAgentAPI | number, Name: string): Observable<DummyAgentAPI> {
    const id = typeof dummyagentdb === 'number' ? dummyagentdb : dummyagentdb.ID;
    const url = `${this.dummyagentsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<DummyAgentAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted dummyagentdb id=${id}`)),
      catchError(this.handleError<DummyAgentAPI>('deleteDummyAgent'))
    );
  }

  // updateFront copy dummyagent to a version with encoded pointers and update to the back
  updateFront(dummyagent: DummyAgent, Name: string): Observable<DummyAgentAPI> {
    let dummyagentAPI = new DummyAgentAPI
    CopyDummyAgentToDummyAgentAPI(dummyagent, dummyagentAPI)
    const id = typeof dummyagentAPI === 'number' ? dummyagentAPI : dummyagentAPI.ID
    const url = `${this.dummyagentsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<DummyAgentAPI>(url, dummyagentAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<DummyAgentAPI>('updateDummyAgent'))
    );
  }

  /** PUT: update the dummyagentdb on the server */
  update(dummyagentdb: DummyAgentAPI, Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI> {
    return this.updateDummyAgent(dummyagentdb, Name, frontRepo)
  }
  updateDummyAgent(dummyagentdb: DummyAgentAPI, Name: string, frontRepo: FrontRepo): Observable<DummyAgentAPI> {
    const id = typeof dummyagentdb === 'number' ? dummyagentdb : dummyagentdb.ID;
    const url = `${this.dummyagentsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<DummyAgentAPI>(url, dummyagentdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated dummyagentdb id=${dummyagentdb.ID}`)
      }),
      catchError(this.handleError<DummyAgentAPI>('updateDummyAgent'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in DummyAgentService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("DummyAgentService" + error); // log to console instead

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
