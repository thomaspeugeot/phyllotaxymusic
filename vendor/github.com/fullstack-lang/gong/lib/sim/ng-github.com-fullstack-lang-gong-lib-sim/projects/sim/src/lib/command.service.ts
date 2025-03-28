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

import { CommandAPI } from './command-api'
import { Command, CopyCommandToCommandAPI } from './command'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { EngineAPI } from './engine-api'

@Injectable({
  providedIn: 'root'
})
export class CommandService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CommandServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private commandsUrl: string

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
    this.commandsUrl = origin + '/api/github.com/fullstack-lang/gong/lib/sim/go/v1/commands';
  }

  /** GET commands from the server */
  // gets is more robust to refactoring
  gets(Name: string, frontRepo: FrontRepo): Observable<CommandAPI[]> {
    return this.getCommands(Name, frontRepo)
  }
  getCommands(Name: string, frontRepo: FrontRepo): Observable<CommandAPI[]> {

    let params = new HttpParams().set("Name", Name)

    return this.http.get<CommandAPI[]>(this.commandsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<CommandAPI[]>('getCommands', []))
      );
  }

  /** GET command by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, Name: string, frontRepo: FrontRepo): Observable<CommandAPI> {
    return this.getCommand(id, Name, frontRepo)
  }
  getCommand(id: number, Name: string, frontRepo: FrontRepo): Observable<CommandAPI> {

    let params = new HttpParams().set("Name", Name)

    const url = `${this.commandsUrl}/${id}`;
    return this.http.get<CommandAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched command id=${id}`)),
      catchError(this.handleError<CommandAPI>(`getCommand id=${id}`))
    );
  }

  // postFront copy command to a version with encoded pointers and post to the back
  postFront(command: Command, Name: string): Observable<CommandAPI> {
    let commandAPI = new CommandAPI
    CopyCommandToCommandAPI(command, commandAPI)
    const id = typeof commandAPI === 'number' ? commandAPI : commandAPI.ID
    const url = `${this.commandsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CommandAPI>(url, commandAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CommandAPI>('postCommand'))
    );
  }
  
  /** POST: add a new command to the server */
  post(commanddb: CommandAPI, Name: string, frontRepo: FrontRepo): Observable<CommandAPI> {
    return this.postCommand(commanddb, Name, frontRepo)
  }
  postCommand(commanddb: CommandAPI, Name: string, frontRepo: FrontRepo): Observable<CommandAPI> {

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CommandAPI>(this.commandsUrl, commanddb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted commanddb id=${commanddb.ID}`)
      }),
      catchError(this.handleError<CommandAPI>('postCommand'))
    );
  }

  /** DELETE: delete the commanddb from the server */
  delete(commanddb: CommandAPI | number, Name: string): Observable<CommandAPI> {
    return this.deleteCommand(commanddb, Name)
  }
  deleteCommand(commanddb: CommandAPI | number, Name: string): Observable<CommandAPI> {
    const id = typeof commanddb === 'number' ? commanddb : commanddb.ID;
    const url = `${this.commandsUrl}/${id}`;

    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CommandAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted commanddb id=${id}`)),
      catchError(this.handleError<CommandAPI>('deleteCommand'))
    );
  }

  // updateFront copy command to a version with encoded pointers and update to the back
  updateFront(command: Command, Name: string): Observable<CommandAPI> {
    let commandAPI = new CommandAPI
    CopyCommandToCommandAPI(command, commandAPI)
    const id = typeof commandAPI === 'number' ? commandAPI : commandAPI.ID
    const url = `${this.commandsUrl}/${id}`;
    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<CommandAPI>(url, commandAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CommandAPI>('updateCommand'))
    );
  }

  /** PUT: update the commanddb on the server */
  update(commanddb: CommandAPI, Name: string, frontRepo: FrontRepo): Observable<CommandAPI> {
    return this.updateCommand(commanddb, Name, frontRepo)
  }
  updateCommand(commanddb: CommandAPI, Name: string, frontRepo: FrontRepo): Observable<CommandAPI> {
    const id = typeof commanddb === 'number' ? commanddb : commanddb.ID;
    const url = `${this.commandsUrl}/${id}`;


    let params = new HttpParams().set("Name", Name)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CommandAPI>(url, commanddb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated commanddb id=${commanddb.ID}`)
      }),
      catchError(this.handleError<CommandAPI>('updateCommand'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CommandService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CommandService" + error); // log to console instead

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
