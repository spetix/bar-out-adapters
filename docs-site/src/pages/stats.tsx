import type {ReactNode} from 'react';
import Layout from '@theme/Layout';
import Heading from '@theme/Heading';
import {useEffect, useState} from 'react';
import styles from './stats.module.css';

interface PackageCoverage {
  package: string;
  coverage: string;
}

interface ProjectStats {
  packagesTested: number;
  packagesPassed: number;
  packagesFailed: number;
  packagesSkipped: number;
  coverageReportUrl: string;
  rawTestResultsUrl: string;
  generatedAt: string;
  packageCoverages: PackageCoverage[];
}

interface GoTestRecord {
  Time: string;
  Action: string;
  Package: string;
  Test?: string;
  Output?: string;
  Elapsed?: number;
}

interface PackageTestSummary {
  packageName: string;
  testsRun: number;
  passed: number;
  failed: number;
  skipped: number;
  coverage?: string;
}

interface TestResultsSummary {
  packages: number;
  testsRun: number;
  passed: number;
  failed: number;
  skipped: number;
  packagesSummary: PackageTestSummary[];
}

function parseGoTestRecords(text: string): GoTestRecord[] {
  return text
    .split('\n')
    .map((line) => line.trim())
    .filter(Boolean)
    .map((line) => {
      try {
        return JSON.parse(line) as GoTestRecord;
      } catch {
        return null;
      }
    })
    .filter((record): record is GoTestRecord => record !== null);
}

function summarizeTestResults(records: GoTestRecord[]): TestResultsSummary {
  const packages = new Map<string, PackageTestSummary>();

  for (const record of records) {
    const pkg = record.Package;
    if (!packages.has(pkg)) {
      packages.set(pkg, {
        packageName: pkg,
        testsRun: 0,
        passed: 0,
        failed: 0,
        skipped: 0,
      });
    }
    const summary = packages.get(pkg)!;

    if (record.Test) {
      if (record.Action === 'pass') {
        summary.passed += 1;
        summary.testsRun += 1;
      } else if (record.Action === 'fail') {
        summary.failed += 1;
        summary.testsRun += 1;
      } else if (record.Action === 'skip') {
        summary.skipped += 1;
        summary.testsRun += 1;
      }
    } else if (record.Action === 'output' && record.Output?.includes('coverage:')) {
      summary.coverage = record.Output.trim();
    }
  }

  const packagesSummary = Array.from(packages.values());
  return {
    packages: packagesSummary.length,
    testsRun: packagesSummary.reduce((sum, pkg) => sum + pkg.testsRun, 0),
    passed: packagesSummary.reduce((sum, pkg) => sum + pkg.passed, 0),
    failed: packagesSummary.reduce((sum, pkg) => sum + pkg.failed, 0),
    skipped: packagesSummary.reduce((sum, pkg) => sum + pkg.skipped, 0),
    packagesSummary,
  };
}

export default function Stats(): ReactNode {
  const [stats, setStats] = useState<ProjectStats | null>(null);
  const [testResults, setTestResults] = useState<TestResultsSummary | null>(null);
  const [loading, setLoading] = useState(true);
  const [loadingTests, setLoadingTests] = useState(true);

  useEffect(() => {
    fetch('/project-stats.json')
      .then((res) => {
        if (!res.ok) {
          throw new Error('Stats not available');
        }
        return res.json();
      })
      .then((data) => {
        setStats(data);
      })
      .catch(() => {
        setStats(null);
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  useEffect(() => {
    fetch('/test-results.json')
      .then((res) => {
        if (!res.ok) {
          throw new Error('Test results not available');
        }
        return res.text();
      })
      .then((text) => {
        const records = parseGoTestRecords(text);
        setTestResults(summarizeTestResults(records));
      })
      .catch(() => {
        setTestResults(null);
      })
      .finally(() => {
        setLoadingTests(false);
      });
  }, []);

  return (
    <Layout
      title="Project Stats"
      description="Project test and coverage statistics for Bar Out Adapters">
      <main className={styles.container}>
        <div className="container">
          <Heading as="h1">Project Statistics</Heading>
          <p className={styles.subtitle}>
            Access the latest test and coverage results for this project.
          </p>

          <div className={styles.actions}>
            <a className="button button--secondary" href="/coverage/index.html">
              View Coverage Report
            </a>
            <a className="button button--secondary" href="/test-results.json">
              View Raw Test Results
            </a>
          </div>

          {loading ? (
            <p className={styles.loading}>Loading stats...</p>
          ) : stats ? (
            <>
              <div className={styles.summaryGrid}>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Packages tested</span>
                  <strong>{stats.packagesTested}</strong>
                </div>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Passed</span>
                  <strong>{stats.packagesPassed}</strong>
                </div>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Failed</span>
                  <strong>{stats.packagesFailed}</strong>
                </div>
                <div className={styles.statCard}>
                  <span className={styles.statLabel}>Skipped</span>
                  <strong>{stats.packagesSkipped}</strong>
                </div>
              </div>

              <div className={styles.tableContainer}>
                <Heading as="h2">Package coverage</Heading>
                {stats.packageCoverages.length > 0 ? (
                  <table className={styles.coverageTable}>
                    <thead>
                      <tr>
                        <th>Package</th>
                        <th>Coverage</th>
                      </tr>
                    </thead>
                    <tbody>
                      {stats.packageCoverages.map((item) => (
                        <tr key={item.package}>
                          <td>{item.package}</td>
                          <td>{item.coverage}</td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                ) : (
                  <p>No coverage data available.</p>
                )}
              </div>

              <div className={styles.tableContainer}>
                <Heading as="h2">Test results</Heading>
                {loadingTests ? (
                  <p>Loading test results...</p>
                ) : testResults ? (
                  <>
                    <div className={styles.summaryGrid}>
                      <div className={styles.statCard}>
                        <span className={styles.statLabel}>Packages</span>
                        <strong>{testResults.packages}</strong>
                      </div>
                      <div className={styles.statCard}>
                        <span className={styles.statLabel}>Tests run</span>
                        <strong>{testResults.testsRun}</strong>
                      </div>
                      <div className={styles.statCard}>
                        <span className={styles.statLabel}>Passed</span>
                        <strong>{testResults.passed}</strong>
                      </div>
                      <div className={styles.statCard}>
                        <span className={styles.statLabel}>Failed</span>
                        <strong>{testResults.failed}</strong>
                      </div>
                      <div className={styles.statCard}>
                        <span className={styles.statLabel}>Skipped</span>
                        <strong>{testResults.skipped}</strong>
                      </div>
                    </div>
                    {testResults.packagesSummary.length > 0 ? (
                      <table className={styles.coverageTable}>
                        <thead>
                          <tr>
                            <th>Package</th>
                            <th>Tests</th>
                            <th>Passed</th>
                            <th>Failed</th>
                            <th>Skipped</th>
                            <th>Coverage</th>
                          </tr>
                        </thead>
                        <tbody>
                          {testResults.packagesSummary.map((pkg) => (
                            <tr key={pkg.packageName}>
                              <td>{pkg.packageName}</td>
                              <td>{pkg.testsRun}</td>
                              <td>{pkg.passed}</td>
                              <td>{pkg.failed}</td>
                              <td>{pkg.skipped}</td>
                              <td>{pkg.coverage ?? 'N/A'}</td>
                            </tr>
                          ))}
                        </tbody>
                      </table>
                    ) : (
                      <p>No package test results available.</p>
                    )}
                  </>
                ) : (
                  <p>Test results are not available.</p>
                )}
              </div>

              <div className={styles.generatedAt}>
                Generated: {stats.generatedAt}
              </div>
            </>
          ) : (
            <div className={styles.fallback}>
              <p>Project stats are not yet available. Check back after the next docs build.</p>
            </div>
          )}
        </div>
      </main>
    </Layout>
  );
}
